package controller

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/usecase"
	"employeeleave/utils/common"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type HistoryController struct {
	router *gin.Engine
	historyUC usecase.HistoryUseCase
}

func (h *HistoryController) createHandler(c *gin.Context) {
	// var history model.HistoryLeave
	var historyRequest dto.HistoryResponseDto
	if err := c.ShouldBindJSON(&historyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var newHistory model.HistoryLeave
	historyRequest.Id = common.GenerateID()
	newHistory.Id = historyRequest.Id
	// newHistory.EmployeeId = historyRequest.EmployeeId
	// newHistory.TransactionId = historyRequest.TransactionId
	newHistory.Employee.ID = historyRequest.EmployeeId
	newHistory.Transaction.ID = historyRequest.TransactionId
	newHistory.DateStart = time.Now()
	newHistory.DateEnd = time.Now()
	newHistory.LeaveDuration = historyRequest.LeaveDuration
	newHistory.StatusLeave = historyRequest.StatusLeave
	if err := h.historyUC.RegisterNewHistory(newHistory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": newHistory})
		return
	}

	// historyResponse := map[string]any{
	// 	"id":       history.Id,
	// 	"employee_id": history.EmployeeId,
	// 	"transaction_id": history.TransactionId,
	// 	// "date_start": history.DateStart,
	// 	// "date_end": history.DateEnd,
	// 	"leave_duration": history.LeaveDuration,
	// 	"status_leave": history.StatusLeave,
	// }

	c.JSON(http.StatusOK, historyRequest)
}

func (h *HistoryController) getHandler(c *gin.Context) {
	id := c.Param("id")
	history, err := h.historyUC.FindHistoryById(id)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get By Id Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   history,
	})
}

func (h *HistoryController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	histories, paging, err := h.historyUC.FindAllHistory(paginationParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Successfully Get All Data",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   histories,
		"paging": paging,
	})
}

// func (h *HistoryController) listHandler(c *gin.Context) {
// 	histories, err := h.historyUC.FindAllHistory()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": histories,
// 	})
// }

func NewHistoryController(r *gin.Engine, usecase usecase.HistoryUseCase) *HistoryController {
	controller := HistoryController{
		router: r,
		historyUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/histories", controller.createHandler)
	rg.GET("/histories/:id", controller.getHandler)
	rg.GET("/histories", controller.listHandler)
	return &controller
}
