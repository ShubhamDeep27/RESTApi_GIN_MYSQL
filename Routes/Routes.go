package routes

import (
	"rest/gin/controller"
	"rest/gin/dao"

	"rest/gin/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	empDao := dao.NewEmployeeDaoImpl()
	empService := service.NewEmployeeServiceImpl(empDao)
	empContoller := controller.NewEmployeeController(empService)

	r := gin.Default()

	r.GET("Employee", empContoller.GetEmployess)
	r.POST("Employee", empContoller.CreateEmployees)
	r.POST("PushData", empContoller.CreateEmployeesImaginary)
	r.PUT("Employee/:id", empContoller.UpdateEmployee)
	r.GET("Employee/:id", empContoller.GetEmployeeById)

	return r
}
