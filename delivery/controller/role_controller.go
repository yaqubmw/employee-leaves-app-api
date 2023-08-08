package controller

import (
	"employeeleave/model"
	"employeeleave/usecase"
	"employeeleave/utils/common"
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

	role.Id = common.GenerateID()
	if err := r.roleUC.RegisterNewRole(role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	roleResponse := map[string]any{
		"id":       role.Id,
		"role_name": role.RoleName,
	}

	c.JSON(http.StatusOK, roleResponse)
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

func NewUserController(r *gin.Engine, usecase usecase.RoleUseCase) *RoleController {
	controller := RoleController{
		router: r,
		roleUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/roles", controller.createHandler)
	rg.GET("/roles", controller.listHandler)
	return &controller
}
