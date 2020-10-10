package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorld just send simple text text on /api route
func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}
