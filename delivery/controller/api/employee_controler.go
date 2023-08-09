package api

import (
	"employeeleave/model"
	"employeeleave/usecase"
	"employeeleave/utils/common"

	"github.com/gin-gonic/gin"
)

type EmplController struct {
	emplUC usecase.EmplUseCase
	router *gin.Engine
}

func (e *EmplController) createHandler(c *gin.Context) {
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

func (e *EmplController) listHandler(c *gin.Context) {
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

func (e *EmplController) getHandler(c *gin.Context) {
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

func (e *EmplController) updateHandler(c *gin.Context) {
	var empl model.Employee
	if err := c.ShouldBindJSON(&empl); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	if err := e.emplUC.UpdateEmpl(empl); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, empl)
}

func (e *EmplController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := e.emplUC.DeleteEmpl(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.String(204, "")
}

func NewEmplController(usecase usecase.EmplUseCase, r *gin.Engine) *EmplController {
	controller := EmplController{
		router: r,
		emplUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/employee", controller.createHandler)
	rg.GET("/employee", controller.listHandler)
	rg.GET("/employee/:id", controller.getHandler)
	rg.PUT("/employee", controller.updateHandler)
	rg.DELETE("/employee/:id", controller.deleteHandler)
	return &controller
}
