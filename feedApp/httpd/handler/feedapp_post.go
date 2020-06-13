package handler

import (
	feedapp "feedApp/httpd/plataform/feedApp"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FeedAppPost handle POST on feedapp
func FeedAppPost(feed *feedapp.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found",
		})
	}

}
