package controllers

import (
	"net/http"
	"sports_team_manager/auth"
	"sports_team_manager/models"
	"sports_team_manager/storage"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterUserRequest struct {
	FirstName   string    `json:"FirstName"`
	LastName    string    `json:"LastName"`
	Username    string    `json:"Username"`
	DOB         time.Time `json:"DOB"`
	Email       string    `json:"Email"`
	PhoneNumber string    `json:"PhoneNumber"`
	Password    string    `json:"Password"`
	Role        string    `json:"Role"`
	Gender      string    `json:"Gender"`
}
type LoginUserRequest struct {
	EmailOrUsername string `json:"EmailOrUsername"`
	Password        string `json:"Password"`
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
	_, err = storage.DB.Exec(`INSERT INTO registered_users (username,first_name, last_name, email, password_hash, phone_number, gender, date_of_birth, role) VALUES ($1, $2,$3,$4,$5,$6,$7,$8,$9);`,
		req.Username, req.FirstName, req.LastName, req.Email, hash, req.PhoneNumber, req.Gender, req.DOB, req.Role)
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
	var userID int
	err := storage.DB.QueryRow(`SELECT user_id, password_hash FROM registered_users WHERE LOWER(email) = LOWER($1) or LOWER(username) =LOWER($1);`, req.EmailOrUsername).Scan(&userID, &storedHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil || !models.CheckPasswordHash(req.Password, storedHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	accessToken, err := auth.GenerateAccessToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation faield"})
		return
	}
	refreshToken, err := auth.GenerateRefreshToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation faield"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful",
		"accessToken":  accessToken,
		"refreshToken": refreshToken})
}
