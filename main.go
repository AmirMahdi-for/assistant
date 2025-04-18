package main

import (
	"assistant/config"
	"assistant/models"
	"assistant/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Load configuration
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// check database connection
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//migrate
	if err := db.AutoMigrate(&models.Message{}); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Load all routes
	routes.SetupRoutes(router)

	// Start server
	if err := router.Run(":7007"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
