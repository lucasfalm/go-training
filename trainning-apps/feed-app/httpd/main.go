package main

import (
	"feedApp/httpd/handler"
	feedapp "feedApp/httpd/plataform/feedApp"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := feedapp.New()
	r := gin.Default() // instance of gin, with middlwares

	// create new route with context
	r.GET("/ping", handler.PingGet())
	r.GET("/feedapp", handler.FeedAppGet(feed))
	r.POST("/feedapp", handler.FeedAppPost(feed))

	r.Run() // prevent close application, keep server listening
}
