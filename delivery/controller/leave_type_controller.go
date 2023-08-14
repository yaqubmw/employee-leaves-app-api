package controller

import (
	"employeeleave/delivery/middleware"
	"employeeleave/model"
	"employeeleave/usecase"
	"employeeleave/utils/common"

	"github.com/gin-gonic/gin"
)

type LeaveTypeController struct {
	leaveTypeUC usecase.LeaveTypeUseCase
	router      *gin.Engine
}

func (lt *LeaveTypeController) createHandler(c *gin.Context) {
	var leavetype model.LeaveType
	if err := c.ShouldBindJSON(&leavetype); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	leavetype.ID = common.GenerateID()
	if err := lt.leaveTypeUC.RegisterNewLeaveType(leavetype); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(201, leavetype)
}

func (lt *LeaveTypeController) listHandler(c *gin.Context) {
	leavetypes, err := lt.leaveTypeUC.FindAllLeaveType()
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   leavetypes,
	})
}

func (lt *LeaveTypeController) getHandler(c *gin.Context) {
	id := c.Param("id")
	leavetype, err := lt.leaveTypeUC.FindByIdLeaveType(id)
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
		"data":   leavetype,
	})
}

func (lt *LeaveTypeController) getByNameHandler(c *gin.Context) {
	name := c.Param("leave_type_name")
	leavetype, err := lt.leaveTypeUC.GetByName(name)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get By Name Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   leavetype,
	})
}

func (lt *LeaveTypeController) updateHandler(c *gin.Context) {
	var leavetype model.LeaveType
	if err := c.ShouldBindJSON(&leavetype); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	if err := lt.leaveTypeUC.UpdateLeaveType(leavetype); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, leavetype)
}

func (lt *LeaveTypeController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := lt.leaveTypeUC.DeleteLeaveType(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.String(204, "")
}

func NewLeaveTypeController(usecase usecase.LeaveTypeUseCase, r *gin.Engine) *LeaveTypeController {
	controller := LeaveTypeController{
		router:      r,
		leaveTypeUC: usecase,
	}
	// daftarkan semua url path disini
	// /leavetype -> GET, POST, PUT, DELETE
	rg := r.Group("/api/v1")
	rg.POST("/leavetypes", middleware.AuthMiddleware(), controller.createHandler)
	rg.GET("/leavetypes", middleware.AuthMiddleware(), controller.listHandler)
	rg.GET("/leavetypes/id/:id", middleware.AuthMiddleware(), controller.getHandler)
	rg.GET("/leavetypes/name/:leave_type_name", middleware.AuthMiddleware(), controller.getByNameHandler)
	rg.PUT("/leavetypes", middleware.AuthMiddleware(), controller.updateHandler)
	rg.DELETE("/leavetypes/id/:id", middleware.AuthMiddleware(), controller.deleteHandler)
	return &controller
}
