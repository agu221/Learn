package controllers

import (
	"net/http"
	"sports_team_manager/models"
	"sports_team_manager/storage"

	"github.com/gin-gonic/gin"
)

type RegisterUserRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
type LoginUserRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func UserRegisterHandler(c *gin.Context) {
	var req RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := models.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = storage.DB.Exec(`INSERT INTO users (email, password_hash) VALUES ($1, $2);`, req.Email, hash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successful"})
}

func UserLoginHandler(c *gin.Context) {
	var req LoginUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedHash string
	err := storage.DB.QueryRow(`SELECT password_hash FROM users WHERE email = $1;`, req.Email).Scan(&storedHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !models.CheckPasswordHash(req.Password, storedHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
