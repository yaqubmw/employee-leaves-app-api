package controller

import (
	"employeeleave/delivery/middleware"
	"employeeleave/model"
	"employeeleave/usecase"

	"github.com/gin-gonic/gin"
)

type PositionController struct {
	positionUC usecase.PositionUseCase
	router     *gin.Engine
}

func (p *PositionController) createHandler(c *gin.Context) {
	var position model.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	if err := p.positionUC.RegisterNewPosition(position); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(201, position)
}

func (p *PositionController) listHandler(c *gin.Context) {
	positions, err := p.positionUC.FindAllPosition()
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
		"data":   positions,
	})
}

func (p *PositionController) getHandler(c *gin.Context) {
	id := c.Param("id")
	position, err := p.positionUC.FindByIdPosition(id)
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
		"data":   position,
	})
}

func (p *PositionController) updateHandler(c *gin.Context) {
	var position model.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	if err := p.positionUC.UpdatePosition(position); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, position)
}

func (p *PositionController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := p.positionUC.DeletePosition(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.String(204, "")
}

func NewPositionController(usecase usecase.PositionUseCase, r *gin.Engine) *PositionController {
	controller := PositionController{
		router:     r,
		positionUC: usecase,
	}
	// daftarkan semua url path disini
	// /position -> GET, POST, PUT, DELETE
	rg := r.Group("/api/v1/admin")
	// path for admin
	rg.POST("/positions", middleware.AuthMiddleware("1"), controller.createHandler)
	rg.GET("/positions", middleware.AuthMiddleware("1"), controller.listHandler)
	rg.GET("/positions/:id", middleware.AuthMiddleware("1"), controller.getHandler)
	rg.PUT("/positions", middleware.AuthMiddleware("1"), controller.updateHandler)
	rg.DELETE("/positions/:id", middleware.AuthMiddleware("1"), controller.deleteHandler)
	return &controller
}
