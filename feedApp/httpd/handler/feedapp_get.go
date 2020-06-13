package handler

import (
	feedapp "feedApp/httpd/plataform/feedApp"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FeedAppGet handle GET on feedapp
func FeedAppGet(feed *feedapp.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}

}
