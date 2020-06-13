package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingGet handle GET on ping
func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found",
		})
	}

}
