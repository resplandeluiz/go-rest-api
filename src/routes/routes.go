package routes

import (
	"clean-arch-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func Bootstrap() {
	router := gin.Default()
	router.GET("/product", controllers.RequestProducts)
	router.Run("localhost:8080")
}
