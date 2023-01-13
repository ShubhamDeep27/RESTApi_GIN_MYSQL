package Routes

import (
	"rest/gin/Controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.GET("Employee", Controller.GetEmployess)
	r.POST("Employee", Controller.CreateEmployees)
	r.PUT("Employee/:id", Controller.UpdateEmployee)
	r.GET("Employee/:id", Controller.GetEmployeeById)

	return r
}
