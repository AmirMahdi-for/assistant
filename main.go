package main

import (
	"assistant/config"
	"assistant/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Print a config value (for testing)
	fmt.Println("DB Host:", cfg.DBHost)

	// Initialize Gin router
	router := gin.Default()

	// Load all routes
	routes.RegisterRoutes(router)

	// Start server
	if err := router.Run(":7001"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
