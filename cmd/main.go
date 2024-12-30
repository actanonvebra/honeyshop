// @title Honeyshop API
// @version 1.0
// @description This is the API documentation for the Honeyshop honeypot project.
// @termsOfService http://swagger.io/terms/

// @contact.name actanonvebra
// @contact.email ibrahimserhatbulut@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/actanonvebra/honeyshop/internal/db"
	"github.com/actanonvebra/honeyshop/internal/handlers"
	"github.com/actanonvebra/honeyshop/internal/middleware"
	"github.com/actanonvebra/honeyshop/internal/repositories"
	"github.com/actanonvebra/honeyshop/internal/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	_ "github.com/actanonvebra/honeyshop/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

var failedAttempts = make(map[string]int)

func main() {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//mac
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set. Go .env file.")
	}
	db.ConnectMongoDB(mongoURI)
	log.Println("MongoDB connection test completed successfully.")

	userRepo := repositories.NewMongoUserRepo("honeyshop", "user")
	productRepo := repositories.NewMongoProductRepo("honeyshop", "products")
	logRepo := repositories.NewMongoLogRepo("honeyshop", "attack_logs")
	checkoutRepo := repositories.NewMongoCheckoutRepo("honeyshop", "checkout")
	cartRepo := repositories.NewMongoCartRepository("honeyshop", "cart")

	userService := &services.DefaultUserService{Repo: userRepo}
	productService := services.NewProductService(productRepo)
	checkoutService := services.NewCheckoutService(checkoutRepo)
	cartService := services.NewCartService(cartRepo)
	logService := services.NewLogService(logRepo)

	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService, logService)
	checkoutHandler := handlers.NewCheckoutHandler(checkoutService)
	cartHandler := handlers.NewCartHandler(cartService, productService, checkoutService)

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	loginRateLimiter := middleware.RateLimiterMiddleWare(logService)

	e.POST("/login", userHandler.Login, loginRateLimiter)

	e.POST("/register", userHandler.Register)

	e.GET("/products", productHandler.GetProducts)

	e.GET("/products/search", productHandler.SearchProducts)

	e.POST("/products", productHandler.AddProduct)

	e.POST("/checkout", checkoutHandler.Checkout)

	e.POST("/cart/checkout/:userID", cartHandler.CheckoutCart)

	log.Println("Server started at:8080")
	e.Logger.Fatal(e.Start(":8080"))

}
