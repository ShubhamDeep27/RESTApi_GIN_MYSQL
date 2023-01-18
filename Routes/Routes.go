package Routes

import (
	"rest/gin/Controller"
	dao "rest/gin/Dao"
	service "rest/gin/Service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	empDao := dao.NewEmployeeDao()
	empService := service.NewEmployeeService(empDao)
	empContoller := Controller.NewEmployeeController(empService)

	r := gin.Default()
	r.GET("Employee", empContoller.GetEmployess)
	r.POST("Employee", empContoller.CreateEmployees)
	r.PUT("Employee/:id", empContoller.UpdateEmployee)
	r.GET("Employee/:id", empContoller.GetEmployeeById)

	return r
}
