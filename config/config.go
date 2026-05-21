package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config stores application configuration .

type Config struct {
	AppName    string
	AppEnv     string
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

// for loading envirnment values to config s

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("warning: .env file not found")
	}

	return &Config{
		AppName:    os.Getenv("APP_NAME"),
		AppEnv:     os.Getenv("APP_ENV"),
		AppPort:    os.Getenv("APP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
	}
}
