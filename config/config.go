package config

import (
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
