package controller

import (
	"employeeleave/usecase"

	"github.com/gin-gonic/gin"
)

type LeaveApplicationController struct {
	router        *gin.Engine
	transactionUC usecase.LeaveApplicationUseCase
}

func NewLeaveApplicationController(r *gin.Engine, usecase usecase.LeaveApplicationUseCase) *LeaveApplicationController {
	controller := LeaveApplicationController{
		router:        r,
		transactionUC: usecase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/applyleaves")
	rg.GET("/applyleaves")
	return &controller
}
