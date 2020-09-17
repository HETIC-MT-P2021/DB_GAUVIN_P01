package main

import (
	"time"

	"github.com/jasongauvin/DB_GAUVIN_P01/database"

	"github.com/jasongauvin/DB_GAUVIN_P01/router"
)

func main() {
	time.Sleep(10 * time.Second)
	database.InitializeDb()
	router.InitRouter()
}
