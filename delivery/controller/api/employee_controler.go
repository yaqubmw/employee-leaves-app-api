package api

import (
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

func NewEmplController(usecase usecase.EmployeeUseCase, r *gin.Engine) *EmployeeController {
	controller := EmployeeController{
		router: r,
		emplUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/employee", controller.createHandler)
	rg.GET("/employee", controller.listHandler)
	rg.GET("/employee/:id", controller.getHandler)
	return &controller
}
