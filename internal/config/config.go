package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
	HttpPort   string // it will be casted to int in the main function
}

func NewConfig() *Config {
	loadEnv()
	return &Config{
		DbUser:     GetEnv("DB_USER"),
		DbPassword: GetEnv("DB_PASSWORD"),
		DbHost:     GetEnv("DB_HOST"),
		DbPort:     GetEnv("DB_PORT"),
		DbName:     GetEnv("DB_NAME"),
		HttpPort:   GetEnv("PORT"),
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s doesnt exist", key)
	}
	return value
}
