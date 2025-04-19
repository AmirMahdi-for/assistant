package routes

import (
	"assistant/controllers"
	"assistant/middleware"
	"assistant/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	repo := repositories.NewMessageRepository()
	controller := controllers.NewMessageController(repo)

	api := router.Group("/")
	api.Use(middleware.CheckLogin(db))
	api.POST("/store", controller.Store)
}
