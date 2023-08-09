package api

// import (
// 	"employeeleave/model"
// 	"employeeleave/model/dto"
// 	"employeeleave/usecase"
// 	"employeeleave/utils/common"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type EmployeeController struct {
// 	router     *gin.Engine
// 	employeeUC usecase.EmplUseCase
// }

// func (e *EmployeeController) createHandler(c *gin.Context) {
// 	var employeeRequest dto.EmployeeRequestDto
// 	if err := c.ShouldBindJSON(&employeeRequest); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 		return
// 	}

// 	var newEmployee model.Employee
// 	newEmployee.ID = common.GenerateID()
// 	newEmployee.PositionID = employeeRequest.PositionID
// 	newEmployee.ManagerID = employeeRequest.ManagerID
// 	newEmployee.Name = employeeRequest.Name
// 	newEmployee.PhoneNumber = employeeRequest.PhoneNumber
// 	newEmployee.Email = employeeRequest.Email
// 	newEmployee.Address = employeeRequest.Address

// 	if err := e.employeeUC.RegisterNewUom(newEmployee); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, newEmployee)
// }

// func (e *EmployeeController) listHandler(c *gin.Context) {
// 	page, _ := strconv.Atoi(c.Query("page"))
// 	limit, _ := strconv.Atoi(c.Query("limit"))
// 	paginationParam := dto.PaginationParam{
// 		Page:  page,
// 		Limit: limit,
// 	}
// 	employees, paging, err := e.employeeUC.FindAllEmployee(paginationParam)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
// 		return
// 	}
// 	status := map[string]any{
// 		"code":        200,
// 		"description": "Get All Data Successfully",
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": status,
// 		"data":   employees,
// 		"paging": paging,
// 	})
// }

// func (e *EmployeeController) getHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	employee, err := e.employeeUC.FindByIdEmployee(id)
// 	if err != nil {
// 		c.JSON(500, gin.H{"err": err.Error()})
// 		return
// 	}
// 	status := map[string]any{
// 		"code":        200,
// 		"description": "Get By Id Data Successfully",
// 	}
// 	c.JSON(200, gin.H{
// 		"status": status,
// 		"data":   employee,
// 	})
// }
// func (e *EmployeeController) updateHandler(c *gin.Context) {
// 	var employeeRequest dto.EmployeeRequestDto
// 	if err := c.ShouldBindJSON(&employeeRequest); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 		return
// 	}

// 	var newEmployee model.Employee
// 	newEmployee.ID = employeeRequest.ID
// 	newEmployee.PositionID = employeeRequest.PositionID
// 	newEmployee.ManagerID = employeeRequest.ManagerID
// 	newEmployee.Name = employeeRequest.Name
// 	newEmployee.PhoneNumber = employeeRequest.PhoneNumber
// 	newEmployee.Email = employeeRequest.Email
// 	newEmployee.Address = employeeRequest.Address

// 	if err := e.employeeUC.UpdateEmployee(newEmployee); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, employeeRequest)
// }

// func (e *EmployeeController) deleteHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := e.employeeUC.DeleteEmployee(id); err != nil {
// 		c.JSON(500, gin.H{"err": err.Error()})
// 		return
// 	}
// 	c.String(204, "")
// }

// func NewEmployeeController(r *gin.Engine, usecase usecase.EmployeeUseCase) *EmployeeController {
// 	controller := EmployeeController{
// 		router:     r,
// 		employeeUC: usecase,
// 	}
// 	rg := r.Group("/api/v1")
// 	rg.POST("/employees", controller.createHandler)
// 	rg.GET("/employees", controller.listHandler)
// 	rg.GET("/employees/:id", controller.getHandler)
// 	rg.PUT("/employees", controller.updateHandler)
// 	rg.DELETE("/employees/:id", controller.deleteHandler)
// 	return &controller
// }
