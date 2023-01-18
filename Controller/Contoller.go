package Controller

import (
	"log"
	"net/http"
	service "rest/gin/Service"
	"rest/gin/models"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeService *service.EmployeeService
}

func NewEmployeeController(es *service.EmployeeService) *EmployeeController {
	return &EmployeeController{employeeService: es}
}

func (ec *EmployeeController) GetEmployess(c *gin.Context) {
	var employee []models.Employee
	err := ec.employeeService.GetAllEmployees(&employee)
	if err != nil {
		log.Fatal("Error while getting Employees")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}
}
func (ec *EmployeeController) GetEmployeeById(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")

	err := ec.employeeService.GetEmployeeById(&employee, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}
}
func (ec *EmployeeController) CreateEmployees(c *gin.Context) {
	var employee models.Employee

	c.BindJSON(&employee)
	err := ec.employeeService.CreateEmployees(&employee)
	if err != nil {
		log.Fatal("Error while Creating Employee")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {

	var employee models.Employee
	id := c.Param("id")
	c.BindJSON(&employee)
	err := ec.employeeService.UpdateEmployee(&employee, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}

}
