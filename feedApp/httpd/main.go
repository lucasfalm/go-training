package main

import (
	"feedApp/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // instance of gin, with middlwares

	// create new route with context
	r.GET("/ping", handler.PingGet())

	r.Run() // prevent close application, keep server listening
}
