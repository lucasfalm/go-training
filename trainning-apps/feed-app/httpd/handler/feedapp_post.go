package handler

import (
	feedapp "feedApp/httpd/plataform/feedApp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type feedAppPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

// FeedAppPost handle POST on feedapp
func FeedAppPost(feed *feedapp.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := feedAppPostRequest{}
		
		c.Bind(&requestBody)

		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found",
		})
	}

}
