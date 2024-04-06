package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua468/bookstore/handlers"
	"github.com/joshua468/bookstore/utils"
)

func main() {
	// Initialize database
	err := utils.InitDB("bookstore.db")
	if err != nil {
		panic(err)
	}
	defer utils.DB.Close()

	// Initialize Gin router
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")

	// Serve the index.html file for the root path
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// Login route
	router.POST("/login", handlers.Login)

	// Run the server
	router.RunTLS(":8080", "server.crt", "server.key") // HTTPS
}
