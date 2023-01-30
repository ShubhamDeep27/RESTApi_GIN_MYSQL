package routes

import (
	Controller "rest/gin/controller"
	dao "rest/gin/dao"
	service "rest/gin/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	empDao := dao.NewEmployeeDaoImpl()
	empService := service.NewEmployeeServiceImpl(empDao)
	empContoller := Controller.NewEmployeeController(empService)

	r := gin.Default()

	r.GET("Employee", empContoller.GetEmployess)
	r.POST("Employee", empContoller.CreateEmployees)
	r.PUT("Employee/:id", empContoller.UpdateEmployee)
	r.GET("Employee/:id", empContoller.GetEmployeeById)

	return r
}
