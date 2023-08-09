package controller

import (
	"employeeleave/model"
	"employeeleave/usecase"
	"employeeleave/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuotaLeaveController struct {
	router *gin.Engine
	quotaLeaveUC usecase.QuotaLeaveUseCase
}

func (q *QuotaLeaveController) createHandler(c *gin.Context) {
	var quotaLeave model.QuotaLeave
	
	if err := c.ShouldBindJSON(&quotaLeave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return 
	}
	
	quotaLeave.ID = common.GenerateID()
	if err := q.quotaLeaveUC.RegisterNewQuotaLeave(quotaLeave); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, quotaLeave)

}

func (q *QuotaLeaveController) listHandler(c *gin.Context) {
	quotaLeaves, err := q.quotaLeaveUC.FindAllQuotaLeave()
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
		"data":   quotaLeaves,
	})
}

func (q *QuotaLeaveController) getHandler(c *gin.Context) {
	id := c.Param("id")
	quotaLeave, err := q.quotaLeaveUC.FindByIdQuotaLeave(id)
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
		"data":   quotaLeave,
	})
}

func (q *QuotaLeaveController) updateHandler(c *gin.Context) {
	var quotaLeave model.QuotaLeave

	if err := c.ShouldBindJSON(&quotaLeave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := q.quotaLeaveUC.UpdateQuotaLeave(quotaLeave); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, quotaLeave)
}

func (q *QuotaLeaveController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := q.quotaLeaveUC.DeleteQuotaLeave(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.String(204, "")
}

func NewQuotaLeaveController(r *gin.Engine, usecase usecase.QuotaLeaveUseCase) *QuotaLeaveController {
	controller := QuotaLeaveController{
		router:        r,
		quotaLeaveUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/quotaleaves", controller.createHandler)
	rg.GET("/quotaleaves", controller.listHandler)
	rg.GET("/quotaleaves/:id", controller.getHandler)
	rg.PUT("/quotaleaves", controller.updateHandler)
	rg.DELETE("/quotaleaves/:id", controller.deleteHandler)

	return &controller
}