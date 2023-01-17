package Controller

import (
	"fmt"
	"log"
	"net/http"
	"rest/gin/dao"
	"rest/gin/models"

	"github.com/gin-gonic/gin"
)

func GetEmployess(c *gin.Context) {
	var employee []models.Employee
	err := dao.GetAllEmployees(&employee)
	if err != nil {
		log.Fatal("Error while getting Employees")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}
}
func GetEmployeeById(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")

	err := dao.GetEmployeeById(&employee, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}
}
func CreateEmployees(c *gin.Context) {
	var employee models.Employee

	c.BindJSON(&employee)
	err := dao.CreateEmployees(&employee)
	if err != nil {
		log.Fatal("Error while Creating Employee")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func UpdateEmployee(c *gin.Context) {

	var employee models.Employee
	id := c.Param("id")
	fmt.Print(id)
	err := dao.GetEmployeeById(&employee, id)
	if err != nil {
		c.JSON(http.StatusNotFound, employee)
	}

	c.BindJSON(&employee)
	err = dao.UpdateEmployee(&employee, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}

}
