package controller

import (
	"employeeleave/delivery/middleware"
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/usecase"
	"employeeleave/utils/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionLeaveController struct {
	router *gin.Engine
	txUC   usecase.TransactionLeaveUseCase
}

func (tl *TransactionLeaveController) createHandler(c *gin.Context) {
	var trx model.TransactionLeave
	if err := c.ShouldBindJSON(&trx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	trx.ID = common.GenerateID()
	if err := tl.txUC.ApplyLeave(trx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	userResponse := map[string]any{
		"id":              trx.ID,
		"employee_id":     trx.EmployeeID,
		"leave_type_id":   trx.LeaveTypeID,
		"status_leave_id": trx.StatusLeaveID,
		"date_start":      trx.DateStart,
		"date_end":        trx.DateEnd,
		"reason":          trx.Reason,
		"submission_date": trx.SubmissionDate,
	}

	c.JSON(http.StatusOK, userResponse)
}

func (tl *TransactionLeaveController) updateStatusHandler(c *gin.Context) {
	var transactionLeave model.TransactionLeave

	if err := c.ShouldBindJSON(&transactionLeave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := tl.txUC.ApproveOrRejectLeave(transactionLeave); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, transactionLeave)
}

func (t *TransactionLeaveController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	employees, paging, err := t.txUC.FindAllEmpl(paginationParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   employees,
		"paging": paging,
	})
}

func (t *TransactionLeaveController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")
	result, err := t.txUC.FindById(id)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Successfully Get By ID Data",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   result,
	})
}

func (t *TransactionLeaveController) getByEmployeeIdHandler(c *gin.Context) {
	id := c.Param("id")
	result, err := t.txUC.FindByIdEmpl(id)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Successfully Get By ID Data",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   result,
	})
}

func NewTransactionController(r *gin.Engine, usecase usecase.TransactionLeaveUseCase) *TransactionLeaveController {
	controller := TransactionLeaveController{
		router: r,
		txUC:   usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/transaction", middleware.AuthMiddleware("2"), controller.createHandler)
	rg.GET("hc/transaction", middleware.AuthMiddleware("4"), controller.listHandler)
	rg.GET("manager/transaction", middleware.AuthMiddleware("3"), controller.listHandler)
	rg.GET("employee/transaction/:id", middleware.AuthMiddleware("2"), controller.getByEmployeeIdHandler)
	rg.GET("hc/transaction/:id", middleware.AuthMiddleware("4"), controller.getByIdHandler)
	rg.GET("manager/transaction/:id", middleware.AuthMiddleware("3"), controller.getByIdHandler)
	rg.PUT("/transaction/update", middleware.AuthMiddleware("3"), controller.updateStatusHandler)
	// rg.GET("/transactions", controller.listHandler)
	return &controller
}
