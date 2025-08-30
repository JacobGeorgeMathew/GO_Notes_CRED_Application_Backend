package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() *Config  {
	godotenv.Load("../.env")

	return &Config{
		DatabaseURL: buildDataBaseURL(),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	//fmt.Println("Default Value")
	return defaultValue
}

func buildDataBaseURL() string {
	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "password")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	dbname := getEnv("DB_NAME", "notes_app")

	return user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?parseTime=true"
}