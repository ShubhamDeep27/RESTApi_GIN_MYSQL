package Controller

import (
	"net/http"
	service "rest/gin/Service"
	"rest/gin/models"
	"rest/gin/util"

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

		c.JSON(http.StatusOK, models.GetAllEmployeeResponse{Data: employee, Status: "Success", Message: "Successfully! Get All Records."})
	}
}
func (ec *EmployeeController) GetEmployeeById(c *gin.Context) {
	id := c.Param("id")
	employee, err := ec.employeeService.GetEmployeeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{ErrorCode: "501", ErrorMsg: util.GET_ERROR_MSGS["501"]})

	} else {
		c.JSON(http.StatusOK, models.GetEmployeeByIdResponse{Data: employee, Status: "Success", Message: "Successfully! Get the Record"})
	}
}
func (ec *EmployeeController) CreateEmployees(c *gin.Context) {
	var employee models.CreateEmployee

	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorCode: "502", ErrorMsg: util.GET_ERROR_MSGS["502"]})
		return
	}

	err := ec.employeeService.CreateEmployees(&employee)
	if err != (models.ErrorResponse{}) {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, models.CreateEmployeeResponse{Data: &employee, Status: "Success", Message: "Successfully! Record has been added"})
	}
}
func (ec *EmployeeController) CreateEmployeesImaginary(c *gin.Context) {

	response, err := ec.employeeService.CreateEmployeesImaginary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{ErrorCode: "505", ErrorMsg: util.GET_ERROR_MSGS["505"]})
	} else {
		c.JSON(http.StatusOK, response)
	}
}
func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {

	var employee models.CreateEmployee
	id := c.Param("id")
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorCode: "506", ErrorMsg: util.GET_ERROR_MSGS["506"]})
		return
	}
	err := ec.employeeService.UpdateEmployee(&employee, id)
	if err != (models.ErrorResponse{}) {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, models.UpdateEmployeeResponse{Data: &employee, Status: "Success", Message: "Successfully! Record has been added"})
	}

}
