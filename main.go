package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thejunghare/build/controllers"
	"github.com/thejunghare/build/models"
)

func main() {
	// Setting up the server
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	models.ConnectDatabase() // Connect to database

	r.GET("/books", controllers.FindBooks)

	r.Run()
}
