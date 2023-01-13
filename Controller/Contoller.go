package Controller

import (
	"fmt"
	"log"
	"net/http"
	"rest/gin/Models"

	"github.com/gin-gonic/gin"
)

func GetEmployess(c *gin.Context) {
	var employee []Models.Employee

	err := Models.GetAllEmployees(&employee)
	if err != nil {
		log.Fatal("Error while getting Employees")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}
}
func GetEmployeeById(c *gin.Context) {
	var employee Models.Employee
	id := c.Param("id")

	err := Models.GetEmployeeById(&employee, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}
}
func CreateEmployees(c *gin.Context) {
	var employee Models.Employee

	c.BindJSON(&employee)
	err := Models.CreateEmployees(&employee)
	if err != nil {
		log.Fatal("Error while Creating Employee")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func UpdateEmployee(c *gin.Context) {

	var employee Models.Employee
	id := c.Param("id")
	fmt.Print(id)
	err := Models.GetEmployeeById(&employee, id)
	if err != nil {
		c.JSON(http.StatusNotFound, employee)
	}

	c.BindJSON(&employee)
	err = Models.UpdateEmployee(&employee, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}

}
