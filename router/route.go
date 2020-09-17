package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jasongauvin/DB_GAUVIN_P01/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/employee", controllers.ShowEmployee)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	fmt.Println("Listening to port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
