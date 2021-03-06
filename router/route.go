package router

import (
	"github.com/jasongauvin/DB_GAUVIN_P01/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {

	api := router.Group("/api")
	{
		api.GET("/", controllers.HelloWorld)

		v1 := api.Group("/v1")
		{
			v1.GET("/employees", controllers.GetEmployees)
			v1.GET("/offices/:id/employees", controllers.GetEmployeesByOfficeCode)
			v1.GET("/offices/:id", controllers.GetOfficeByCode)
			v1.GET("/customer/:id", controllers.GetCustomerByCode)
			v1.GET("/order/:id", controllers.GetOrderByCode)
		}
	}
}
