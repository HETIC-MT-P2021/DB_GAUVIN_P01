package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/jasongauvin/DB_GAUVIN_P01/database"

	"github.com/jasongauvin/DB_GAUVIN_P01/router"
)

func main() {

	time.Sleep(5 * time.Second)

	database.InitializeDb()

	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "Authorization",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.InitRouter(r)

	log.Fatal(r.Run(":8080")) // listen and serve on 8080
}
