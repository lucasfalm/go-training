package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() // instance of gin, with middlwares

	// create new route with context
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // prevent close application, keep server listening
}
