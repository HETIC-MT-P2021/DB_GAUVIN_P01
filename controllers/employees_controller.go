package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/DB_GAUVIN_P01/repository"
)

func ShowEmployee(c *gin.Context) {
	employee := repository.SelectEmployee()
	fmt.Println(employee)
}
