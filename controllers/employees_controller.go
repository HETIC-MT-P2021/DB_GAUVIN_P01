package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/DB_GAUVIN_P01/models"
	"github.com/jasongauvin/DB_GAUVIN_P01/repository"
)

func ShowEmployees(c *gin.Context) {
	var employees []*models.Employee
	var err error

	employees, err = repository.SelectEmployees()

	fmt.Println(err)
	fmt.Println("employee selected: /n", employees)

	if err != nil {
		c.JSON(http.StatusNotFound, "Could'nt fetch employees.")
		return
	}

	c.JSON(http.StatusOK, employees)
}
