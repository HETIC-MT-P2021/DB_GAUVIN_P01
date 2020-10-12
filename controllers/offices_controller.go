package controllers

import (
	"net/http"

	"github.com/jasongauvin/DB_GAUVIN_P01/models"

	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/DB_GAUVIN_P01/repository"
	"github.com/jasongauvin/DB_GAUVIN_P01/utils"
)

func GetOfficeByCode(c *gin.Context) {
	var offices *models.Office
	var err error
	id := utils.ParseStringToUint64(c.Param("id"))

	offices, err = repository.FindOfficeByCode(id)

	if err != nil {
		c.JSON(http.StatusNotFound, "Couldn't fetch offices by code.")
		return
	}

	c.JSON(http.StatusOK, offices)
}
