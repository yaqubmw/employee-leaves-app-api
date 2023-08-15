package controller

import (
	"employeeleave/delivery/middleware"
	"employeeleave/model"
	"employeeleave/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	router *gin.Engine
	roleUC usecase.RoleUseCase
}

func (r *RoleController) createHandler(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := r.roleUC.RegisterNewRole(role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	roleResponse := map[string]any{
		"id":        role.Id,
		"role_name": role.RoleName,
	}

	c.JSON(http.StatusOK, roleResponse)
}

func (r *RoleController) getHandler(c *gin.Context) {
	id := c.Param("id")
	role, err := r.roleUC.FindByIdRole(id)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get Data By Id Success",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   role,
	})
}

func (r *RoleController) listHandler(c *gin.Context) {
	roles, err := r.roleUC.FindAllRole()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": roles,
	})
}

func (r *RoleController) updateHandler(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	// call usecase
	if err := r.roleUC.UpdateRole(role); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, role)
}

func (r *RoleController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := r.roleUC.DeleteRole(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        204,
		"description": "Delete Data By Id Success",
	}
	c.JSON(204, gin.H{
		"status": status,
	})
}

func NewRoleController(r *gin.Engine, usecase usecase.RoleUseCase) *RoleController {
	controller := RoleController{
		router: r,
		roleUC: usecase,
	}

	rg := r.Group("/api/v1/admin")
	// path for admin
	rg.POST("/roles", middleware.AuthMiddleware("1"), controller.createHandler)
	rg.GET("/roles", middleware.AuthMiddleware("1"), controller.listHandler)
	rg.GET("/roles/:id", middleware.AuthMiddleware("1"), controller.getHandler)
	rg.PUT("/roles", middleware.AuthMiddleware("1"), controller.updateHandler)
	rg.DELETE("/roles/:id", middleware.AuthMiddleware("1"), controller.deleteHandler)
	return &controller
}
