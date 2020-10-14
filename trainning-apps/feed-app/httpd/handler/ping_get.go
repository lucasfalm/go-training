package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingGet handle GET on ping
func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Request.FormValue("i")
		fmt.Printf("params type %T\n", param)

		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found",
		})
	}

}
