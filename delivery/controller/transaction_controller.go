package controller

import (
	"employeeleave/model"
	"employeeleave/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionLeaveController struct {
	router        *gin.Engine
	transactionUC usecase.TransactionLeaveUseCase
}

func (tl *TransactionLeaveController) updateStatusHandler(c *gin.Context) {
	var transactionLeave model.TransactionLeave

	if err := c.ShouldBindJSON(&transactionLeave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := tl.transactionUC.ApproveOrRejectLeave(transactionLeave); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, transactionLeave)
}

func NewTransactionLeaveController(r *gin.Engine, usecase usecase.TransactionLeaveUseCase) *TransactionLeaveController {
	controller := TransactionLeaveController{
		router:        r,
		transactionUC: usecase,
	}
	rg := r.Group("/api/v1")
	rg.PUT("/decisions", controller.updateStatusHandler)
	return &controller
}
