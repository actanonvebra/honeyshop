package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/actanonvebra/honeyshop/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set. Go .env file.")
	}
	db.ConnectMongoDB(mongoURI)
	log.Println("MongoDB connection test completed successfully.")

}
