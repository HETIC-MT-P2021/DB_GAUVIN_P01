package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/DB_GAUVIN_P01/models"
	"github.com/jasongauvin/DB_GAUVIN_P01/repository"
	"github.com/jasongauvin/DB_GAUVIN_P01/utils"
)

func GetCustomerByCode(c *gin.Context) {
	var customer *models.Customer
	var err error
	id := utils.ParseStringToUint64(c.Param("id"))

	customer, err = repository.FindCustomerByCode(id)

	if err != nil {
		c.JSON(http.StatusNotFound, "Couldn't fetch customers by code.")
		return
	}

	c.JSON(http.StatusOK, customer)
}
