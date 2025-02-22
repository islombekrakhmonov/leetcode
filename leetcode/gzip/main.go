package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Use gzip middleware
	r.Use(gzip.Gzip(gzip.BestSpeed))

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // Start the server on port 8080
}
