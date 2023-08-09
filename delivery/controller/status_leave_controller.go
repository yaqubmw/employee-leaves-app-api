package controller

import (
	"employeeleave/model"
	"employeeleave/usecase"
	"employeeleave/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusLeaveController struct {
	router *gin.Engine
	statusLeaveUC usecase.StatusLeaveUseCase
}

func (s *StatusLeaveController) createHandler(c *gin.Context) {
	var statusLeave model.StatusLeave
	
	if err := c.ShouldBindJSON(&statusLeave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return 
	}
	
	statusLeave.ID = common.GenerateID()
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
	rg.POST("/statusleaves", controller.createHandler)
	rg.GET("/statusleaves", controller.listHandler)
	rg.GET("/statusleaves/:id", controller.getHandler)
	rg.PUT("/statusleaves", controller.updateHandler)
	rg.DELETE("/statusleaves/:id", controller.deleteHandler)

	return &controller
}