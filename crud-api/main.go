package main

import (
	"context"
	"crud-api/config"
	"crud-api/controllers"
	"crud-api/middleware"
	"crud-api/repository"
	"crud-api/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		panic("Failed to connect to MongoDB")
	}
	db := client.Database("Go-Crud-One")

	// Initialize repositories and controllers
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	locationRepo := repository.NewLocationRepository(db)

	userController := controllers.UserController{Repo: userRepo}
	authController := controllers.NewAuthController(userRepo)
	productController := controllers.NewProductController(productRepo)
	locationController := controllers.NewLocationController(locationRepo)

	// Setup routes
	r := gin.Default()
	r.Use(middleware.Logger())
	routes.SetupRoutes(r, &userController, authController, productController, locationController)

	// Run server
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server")
	}
}
