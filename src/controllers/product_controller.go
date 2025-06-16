package controllers

import (
	"net/http"

	usecaseproduct "clean-arch-api/src/usecase/product"

	"github.com/gin-gonic/gin"
)

func RequestProducts(c *gin.Context) {

	product := usecaseproduct.RegisterProduct()

	c.JSON(http.StatusOK, product)
}
