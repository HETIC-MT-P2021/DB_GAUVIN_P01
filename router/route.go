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
			v1.GET("/employee", controllers.ShowEmployee)
			// v1.GET("/customers/:number", controllers.GetCustomerByNumber)
			// v1.GET("/customers/:number/orders", controllers.GetOrdersByCustomerNumber)
			// v1.GET("/orders/:number", controllers.GetOrderByNumber)
			// v1.GET("/orders/:number/details", controllers.GetOrderDetailsByOrderNumber)
			// v1.GET("/offices/:code/employees", controllers.GetEmployeesByOfficeCode)
			// v1.GET("/offices/:code", controllers.GetOfficeByCode)
		}
	}
}
