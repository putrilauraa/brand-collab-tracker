package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Attention: Failed to load .env file. Using Environment Variable system.")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("FATAL: JWT_SECRET environment variable not found!")
	}
	return secret
}