package routes

import (
	"assistant/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/store", controllers.Store)
}
