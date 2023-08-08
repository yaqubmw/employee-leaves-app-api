package api

import (
	"employeeleave/model"
	"employeeleave/model/dto"
	"employeeleave/usecase"
	"employeeleave/utils/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	router  *gin.Engine
	usecase usecase.EmployeeUseCase
}

func (e *EmployeeController) createHandler(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	employee.ID = common.GenerateID()
	if err := e.usecase.RegisterNewEmployee(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, employee)

	employee.ManagerID = common.GenerateID()
	if err := e.usecase.RegisterNewEmployee(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, employee)

}
func (e *EmployeeController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	employees, paging, err := e.usecase.FindAllEmployee(paginationParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   employees,
		"paging": paging,
	})
}
func (e *EmployeeController) getHandler(c *gin.Context) {
	id := c.Param("id")
	employee, err := e.usecase.FindByIdEmployee(id)
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
		"data":   employee,
	})
}
func (e *EmployeeController) updateHandler(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := e.usecase.UpdateEmployee(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}
func (e *EmployeeController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := e.usecase.DeleteEmployee(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.String(204, "")
}

func NewEmployeeController(r *gin.Engine, usecase usecase.EmployeeUseCase) *EmployeeController {
	controller := EmployeeController{
		router:  r,
		usecase: usecase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/employees", controller.createHandler)
	rg.GET("/employees", controller.listHandler)
	rg.GET("/employees/:id", controller.getHandler)
	rg.PUT("/employees", controller.updateHandler)
	rg.DELETE("/employees/:id", controller.deleteHandler)
	return &controller
}
