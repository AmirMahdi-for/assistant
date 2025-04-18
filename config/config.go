package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct for holding all environment variables
type Config struct {
	DBConnection         string
	DBHost               string
	DBPort               string
	DBDatabase           string
	DBUsername           string
	DBPassword           string
	KafkaBrokers         string
	KafkaConsumerGroupID string
	TopicSalesBot        string
	SalesBotGroupID      string
	TopicTryOn           string
	TryOnGroupID         string
	TopicProductSearch   string
	ProductSearchGroupID string
	TopicWebSearch       string
	WebSearchGroupID     string
	TopicBackend         string
	BackendGroupID       string
}

// LoadConfig loads environment variables from .env file
func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Create a config instance to hold the values
	config := &Config{
		DBConnection:         os.Getenv("DB_CONNECTION"),
		DBHost:               os.Getenv("DB_HOST"),
		DBPort:               os.Getenv("DB_PORT"),
		DBDatabase:           os.Getenv("DB_DATABASE"),
		DBUsername:           os.Getenv("DB_USERNAME"),
		DBPassword:           os.Getenv("DB_PASSWORD"),
		KafkaBrokers:         os.Getenv("KAFKA_BROKERS"),
		KafkaConsumerGroupID: os.Getenv("KAFKA_CONSUMER_GROUP_ID"),
		TopicSalesBot:        os.Getenv("TOPIC_SALES_BOT"),
		SalesBotGroupID:      os.Getenv("SALES_BOT_GROUP_ID"),
		TopicTryOn:           os.Getenv("TOPIC_TRY_ON"),
		TryOnGroupID:         os.Getenv("TRY_ON_GROUP_ID"),
		TopicProductSearch:   os.Getenv("TOPIC_PRODUCT_SEARCH"),
		ProductSearchGroupID: os.Getenv("PRODUCT_SEARCH_GROUP_ID"),
		TopicWebSearch:       os.Getenv("TOPIC_WEB_SEARCH"),
		WebSearchGroupID:     os.Getenv("WEB_SEARCH_GROUP_ID"),
		TopicBackend:         os.Getenv("TOPIC_BACKEND"),
		BackendGroupID:       os.Getenv("BACKEND_GROUP_ID"),
	}

	return config, nil
}

// ConnectDatabase establishes a connection to the database using environment variables
func ConnectDatabase() (*gorm.DB, error) {
	// Load config
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Set up the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBDatabase)

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connection successful")
	return db, nil
}
