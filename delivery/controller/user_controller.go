package controller

import (
	"employeeleave/delivery/middleware"
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/usecase"
	"employeeleave/utils/common"
	"net/http"
	"strconv"

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
		"id":        user.ID,
		"username":  user.Username,
		"role_id":   user.RoleId,
		"is_active": user.IsActive,
	}

	c.JSON(http.StatusOK, userResponse)
}

func (u *UserController) updateHandler(c *gin.Context) {
	var user model.UserCredential
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := u.userUC.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	userResponse := map[string]any{
		"id":        user.ID,
		"username":  user.Username,
		"role_id":   user.RoleId,
		"is_active": user.IsActive,
	}

	c.JSON(http.StatusOK, userResponse)
}

func (u *UserController) getHandler(c *gin.Context) {
	id := c.Param("id")
	user, err := u.userUC.FindByIdUser(id)
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
		"data":   user,
	})
}

func (u *UserController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	users, paging, err := u.userUC.FindAllUser(paginationParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Successfully Get All Data",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   users,
		"paging": paging,
	})
}

func NewUserController(r *gin.Engine, usecase usecase.UserUseCase) *UserController {
	controller := UserController{
		router: r,
		userUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/users", controller.createHandler)
	rg.GET("/users", middleware.AuthMiddleware("1"), controller.listHandler)
	rg.GET("/users/:id", middleware.AuthMiddleware("1"), controller.getHandler)
	rg.PUT("/users", middleware.AuthMiddleware("1"), controller.updateHandler)
	return &controller
}
