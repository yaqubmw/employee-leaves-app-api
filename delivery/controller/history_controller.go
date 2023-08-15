package controller

import (
	"employeeleave/delivery/middleware"
	"employeeleave/model/dto"
	"employeeleave/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HistoryController struct {
	router    *gin.Engine
	historyUC usecase.HistoryUseCase
}

func (h *HistoryController) getHandler(c *gin.Context) {
	id := c.Param("id")
	history, err := h.historyUC.FindHistoryById(id)
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
		"data":   history,
	})
}

func (h *HistoryController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	histories, paging, err := h.historyUC.FindAllHistory(paginationParam)
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
		"data":   histories,
		"paging": paging,
	})
}

func NewHistoryController(r *gin.Engine, usecase usecase.HistoryUseCase) *HistoryController {
	controller := HistoryController{
		router:    r,
		historyUC: usecase,
	}

	rg := r.Group("/api/v1/admin")
	// path for admin
	rg.GET("/histories/:id", middleware.AuthMiddleware("1"), controller.getHandler)
	rg.GET("/histories", middleware.AuthMiddleware("1"), controller.listHandler)
	return &controller
}
