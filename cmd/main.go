package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/actanonvebra/honeyshop/internal/db"
	"github.com/actanonvebra/honeyshop/internal/handlers"
	"github.com/actanonvebra/honeyshop/internal/repositories"
	"github.com/actanonvebra/honeyshop/internal/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	userRepo := repositories.NewMongoUserRepo("honeyshop", "user")
	userService := &services.DefaultUserService{Repo: userRepo}
	userHandler := handlers.NewUserHandler(userService)

	productRepo := repositories.NewMongoProductRepo("honeyshop", "products")
	productService := &services.DefaultProductService{Repo: productRepo}

	logRepo := repositories.NewMongoLogRepo("honeyshop", "attack_logs")
	LogService := services.NewLogService(logRepo)
	productHandler := handlers.NewProductHandler(productService, LogService)

	e := echo.New()

	e.POST("/login", userHandler.Login)
	e.POST("/register", userHandler.Register)
	e.GET("/products", productHandler.GetProducts)
	e.GET("/products/search", productHandler.SearchProducts)
	log.Println("Server started at:8080")
	e.Logger.Fatal(e.Start(":8080"))

}
