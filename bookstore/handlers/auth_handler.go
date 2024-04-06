package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/bookstore/models"
	"github.com/joshua468/bookstore/utils"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	// Query the database for user credentials
	row := utils.DB.QueryRow("SELECT id, username, password, email FROM users WHERE username=?", user.Username)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare passwords
	if err := utils.ComparePasswords(user.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Return user data upon successful login
	c.JSON(http.StatusOK, gin.H{"user": user})
}
