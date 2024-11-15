package routes

import (
	"crud-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/:id", userController.GetUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}
}
