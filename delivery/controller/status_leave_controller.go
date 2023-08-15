package controller

import (
	"employeeleave/delivery/middleware"
	"employeeleave/model"
	"employeeleave/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusLeaveController struct {
	router        *gin.Engine
	statusLeaveUC usecase.StatusLeaveUseCase
}

func (s *StatusLeaveController) createHandler(c *gin.Context) {
	var statusLeave model.StatusLeave

	if err := c.ShouldBindJSON(&statusLeave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := s.statusLeaveUC.RegisterNewStatusLeave(statusLeave); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, statusLeave)

}

func (s *StatusLeaveController) listHandler(c *gin.Context) {
	statusLeaves, err := s.statusLeaveUC.FindAllStatusLeave()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}

	c.JSON(200, gin.H{
		"status": status,
		"data":   statusLeaves,
	})
}

func (s *StatusLeaveController) getHandler(c *gin.Context) {
	id := c.Param("id")
	statusLeave, err := s.statusLeaveUC.FindByIdStatusLeave(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	status := map[string]any{
		"code":        200,
		"description": "Get by Id Data Successfully",
	}

	c.JSON(200, gin.H{
		"status": status,
		"data":   statusLeave,
	})
}

func (s *StatusLeaveController) updateHandler(c *gin.Context) {
	var statusLeave model.StatusLeave

	if err := c.ShouldBindJSON(&statusLeave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := s.statusLeaveUC.UpdateStatusLeave(statusLeave); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, statusLeave)
}

func (s *StatusLeaveController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := s.statusLeaveUC.DeleteStatusLeave(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.String(204, "")
}

func NewStatusLeaveController(r *gin.Engine, usecase usecase.StatusLeaveUseCase) *StatusLeaveController {
	controller := StatusLeaveController{
		router:        r,
		statusLeaveUC: usecase,
	}

	rg := r.Group("/api/v1")
	// path for admin
	rg.POST("/admin/statusleaves", middleware.AuthMiddleware("1"), controller.createHandler)
	rg.GET("/admin/statusleaves", middleware.AuthMiddleware("1"), controller.listHandler)
	rg.GET("/admin/statusleaves/:id", middleware.AuthMiddleware("1"), controller.getHandler)
	rg.PUT("/admin/statusleaves", middleware.AuthMiddleware("1"), controller.updateHandler)
	rg.DELETE("/admin/statusleaves/:id", middleware.AuthMiddleware("1"), controller.deleteHandler)
	// path for manager
	rg.GET("/manager/statusleaves", middleware.AuthMiddleware("3"), controller.listHandler)
	rg.GET("/manager/statusleaves/:id", middleware.AuthMiddleware("3"), controller.getHandler)

	return &controller
}
