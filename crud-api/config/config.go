package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var MongoURI string

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	MongoURI = os.Getenv("MONGO_URI")
	if MongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env")
	}
}
