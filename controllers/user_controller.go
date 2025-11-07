package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"brand-collab-tracker/middlewares"
	"brand-collab-tracker/repositories"
)

func RegisterHandler(c *gin.Context) {
	var req repositories.UserAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	user, err := repositories.RegisterUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfull registration", "username": user.Username})
}

func LoginHandler(c *gin.Context) {
	var req repositories.UserAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	user, err := repositories.VerifyUser(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := middlewares.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfull login", "token": tokenString})
}