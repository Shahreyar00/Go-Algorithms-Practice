package routes

import (
	"crud-api/controllers"
	"crud-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController, authController *controllers.AuthController, productController *controllers.ProductController, locationController *controllers.LocationController) {
	r.Use(middleware.RateLimitMiddleware())

	// Auth routes
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/signup", userController.CreateUser)
	}

	// User routes
	userRoutes := r.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.GET("/:id", userController.GetUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	// Product routes
	productRoutes := r.Group("/products")
	productRoutes.Use(middleware.AuthMiddleware())
	{
		productRoutes.POST("/", productController.CreateProduct)
		productRoutes.GET("/:id", productController.GetProduct)
		productRoutes.GET("/", productController.GetProducts)
		productRoutes.GET("/filter", productController.GetProductsByCategoryOrType)
		productRoutes.PUT("/:id", productController.UpdateProduct)
		productRoutes.DELETE("/:id", productController.DeleteProduct)
	}

	// Location routes
	locationRoutes := r.Group("/locations")
	{
		locationRoutes.POST("/", locationController.AddLocation)
		locationRoutes.GET("/nearby", locationController.GetNearbyLocations)
		locationRoutes.GET("/region", locationController.DetermineRegion)
	}
}
