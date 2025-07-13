package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL string
	Port string
	RabbitURL string
	TransactionQueueName string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	cfg := Config{
		DBURL: getEnv("DB_URL", "postgres://postgres:postgres@localhost:5432/gofinance?sslmode=disable"),
		Port:  getEnv("PORT", "8080"),
		RabbitURL: getEnv("RabbitURL", "amqp://guest:guest@localhost:5672/"),
		TransactionQueueName: getEnv("TransactionQueueName","transactions"),
	}

	return cfg
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}