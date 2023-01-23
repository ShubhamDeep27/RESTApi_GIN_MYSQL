package Controller

import (
	"net/http"
	service "rest/gin/Service"
	"rest/gin/models"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeService service.EmployeeService
}

func NewEmployeeController(es service.EmployeeService) *EmployeeController {
	return &EmployeeController{employeeService: es}
}

func (ec *EmployeeController) GetEmployess(c *gin.Context) {

	employee, err := ec.employeeService.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Records not Found"})

	} else {
		c.JSON(http.StatusOK, employee)
	}
}
func (ec *EmployeeController) GetEmployeeById(c *gin.Context) {
	id := c.Param("id")
	employee, err := ec.employeeService.GetEmployeeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Employee Not Exist"})

	} else {
		c.JSON(http.StatusOK, employee)
	}
}
func (ec *EmployeeController) CreateEmployees(c *gin.Context) {
	var employee models.Employee

	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ec.employeeService.CreateEmployees(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating Employee"})
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {

	var employee models.Employee
	id := c.Param("id")
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ec.employeeService.UpdateEmployee(&employee, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Employee Not Exist"})
	} else {
		c.JSON(http.StatusOK, employee)
	}

}
