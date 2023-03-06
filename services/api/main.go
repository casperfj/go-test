package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create gin router
	router := gin.Default()

	// Create a default router group
	root := router.Group("/")

	// Create get route for root
	root.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// Start the server
	router.Run(":5000")
}