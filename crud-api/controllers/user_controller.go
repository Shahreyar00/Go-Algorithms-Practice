package controllers

import (
	"crud-api/models"
	"crud-api/repository"
	"crud-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	Repo *repository.UserRepository
}

// Handlers
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = primitive.NewObjectID()
	if err := ctrl.Repo.CreateUser(user); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}
	utils.RespondSuccess(c, http.StatusCreated, user)
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	user, err := ctrl.Repo.GetUserByID(objID)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondSuccess(c, http.StatusOK, user)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := ctrl.Repo.UpdateUser(objID, user); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update user")
		return
	}
	utils.RespondSuccess(c, http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	if err := ctrl.Repo.DeleteUser(objID); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete user")
		return
	}
	utils.RespondSuccess(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}
