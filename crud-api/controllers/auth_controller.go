package controllers

import (
	"crud-api/repository"
	"crud-api/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserRepo *repository.UserRepository
}

func NewAuthController(userRepo *repository.UserRepository) *AuthController {
	return &AuthController{UserRepo: userRepo}
}

func (a *AuthController) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Authenticate user
	user, err := a.UserRepo.AuthenticateUser(loginData.Email, loginData.Password)
	if err != nil {
		userJSON, _ := json.Marshal(user)
		log.Printf("Error: %v, user: %s, loginPassword: %s", err, string(userJSON), loginData.Password)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.Hex(), "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
