package controller

import (
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

// func (tl *TransactionLeaveController) listHandler(c *gin.Context) {
// 	page, _ := strconv.Atoi(c.Query("page"))
// 	limit, _ := strconv.Atoi(c.Query("limit"))
// 	paginationParam := dto.PaginationParam{
// 		Page:  page,
// 		Limit: limit,
// 	}
// 	users, paging, err := tl.txUC.FindAllUser(paginationParam)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
// 		return
// 	}
// 	status := map[string]any{
// 		"code":        200,
// 		"description": "Successfully Get All Data",
// 	}
// 	c.JSON(200, gin.H{
// 		"status": status,
// 		"data":   users,
// 		"paging": paging,
// 	})
// }

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

func (t *TransactionLeaveController) getHandler(c *gin.Context) {
	id := c.Param("id")
	empl, err := t.txUC.FindByIdEmpl(id)
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
		"data":   empl,
	})
}
func NewTransactionController(r *gin.Engine, usecase usecase.TransactionLeaveUseCase) *TransactionLeaveController {
	controller := TransactionLeaveController{
		router: r,
		txUC:   usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/transaction", controller.createHandler)
	rg.GET("/transaction", controller.listHandler)
	rg.GET("/transaction/:id", controller.getHandler)
	return &controller
}
