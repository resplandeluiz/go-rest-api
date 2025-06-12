package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Luiz Resplande",
	})
}
