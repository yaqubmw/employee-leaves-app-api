package controller

import (
	"employeeleave/model"
	"employeeleave/usecase"
	"employeeleave/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router *gin.Engine
	userUC usecase.UserUseCase
}

func (u *UserController) createHandler(c *gin.Context) {
	var user model.UserCredential
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	user.ID = common.GenerateID()
	if err := u.userUC.RegisterNewUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	userResponse := map[string]any{
		"id":       user.ID,
		"username": user.Username,
		"role_id":  user.RoleId,
	}

	c.JSON(http.StatusOK, userResponse)
}

func (u *UserController) listHandler(c *gin.Context) {
	users, err := u.userUC.FindAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func NewUserController(r *gin.Engine, usecase usecase.UserUseCase) *UserController {
	controller := UserController{
		router: r,
		userUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/users", controller.createHandler)
	rg.GET("/users", controller.listHandler)
	return &controller
}
