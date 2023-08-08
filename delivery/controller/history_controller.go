package controller

import (
	"employeeleave/model"
	"employeeleave/usecase"
	"employeeleave/utils/common"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HistoryController struct {
	router *gin.Engine
	historyUC usecase.HistoryUseCase
}

func (h *HistoryController) createHandler(c *gin.Context) {
	var history model.HistoryLeave
	if err := c.ShouldBindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	history.Id = common.GenerateID()
	history.DateStart = time.Now()
	history.DateEnd = time.Now()
	if err := h.historyUC.RegisterNewHistory(history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	historyResponse := map[string]any{
		"id":       history.Id,
		"employee_id": history.EmployeeId,
		"transaction_id": history.TransactionId,
		// "date_start": history.DateStart,
		// "date_end": history.DateEnd,
		"leave_duration": history.LeaveDuration,
		"StatusLeave": history.StatusLeave,
	}

	c.JSON(http.StatusOK, historyResponse)
}

func (h *HistoryController) listHandler(c *gin.Context) {
	histories, err := h.historyUC.FindAllHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": histories,
	})
}

func NewHistoryController(r *gin.Engine, usecase usecase.HistoryUseCase) *HistoryController {
	controller := HistoryController{
		router: r,
		historyUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/histories", controller.createHandler)
	rg.GET("/histories", controller.listHandler)
	return &controller
}
