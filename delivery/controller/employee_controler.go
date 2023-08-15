package controller

import (
	"employeeleave/delivery/middleware"
	"employeeleave/model"
	"employeeleave/usecase"
	"employeeleave/utils/common"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	emplUC usecase.EmployeeUseCase
	router *gin.Engine
}

func (e *EmployeeController) createHandler(c *gin.Context) {
	var empl model.Employee
	if err := c.ShouldBindJSON(&empl); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	empl.ID = common.GenerateID()
	if err := e.emplUC.RegisterNewEmpl(empl); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(201, empl)
}

func (e *EmployeeController) listHandler(c *gin.Context) {
	empls, err := e.emplUC.FindAllEmpl()
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
		"data":   empls,
	})
}

func (e *EmployeeController) getHandler(c *gin.Context) {
	id := c.Param("id")
	empl, err := e.emplUC.FindByIdEmpl(id)
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

func (e *EmployeeController) updateHandler(c *gin.Context) {
	var employe model.Employee
	if err := c.ShouldBindJSON(&employe); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	if err := e.emplUC.UpdateEmpl(employe); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, employe)
}

func NewEmplController(usecase usecase.EmployeeUseCase, r *gin.Engine) *EmployeeController {
	controller := EmployeeController{
		router: r,
		emplUC: usecase,
	}

	rg := r.Group("/api/v1")
	// path for admin
	rg.POST("/admin/profile", middleware.AuthMiddleware("1"), controller.createHandler)
	rg.GET("/admin/profile", middleware.AuthMiddleware("1"), controller.listHandler)
	rg.GET("/admin/profile/:id", middleware.AuthMiddleware("1"), controller.getHandler)
	rg.PUT("/admin/profile", middleware.AuthMiddleware("1"), controller.updateHandler)
	// path for employee
	rg.GET("/employee/profile/:id", middleware.AuthMiddleware("2"), controller.getHandler)
	rg.PUT("/employee/profile", middleware.AuthMiddleware("2"), controller.updateHandler)
	// path for manager
	rg.GET("/manager/profile/:id", middleware.AuthMiddleware("3"), controller.getHandler)
	rg.PUT("/manager/profile", middleware.AuthMiddleware("3"), controller.updateHandler)
	// path for hc
	rg.GET("/hc/profile/:id", middleware.AuthMiddleware("4"), controller.getHandler)
	rg.PUT("/hc/profile", middleware.AuthMiddleware("4"), controller.updateHandler)
	return &controller
}
