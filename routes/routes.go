package routes

import (
	"assistant/controllers"
	"assistant/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	api := router.Group("/")
	api.Use(middleware.CheckLogin(db))
	api.POST("/store", controllers.Store)
}
