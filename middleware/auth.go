package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/taivama/golang-training/utils"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.String(http.StatusBadRequest, "no authorization header found\n")
			c.Abort()
			return
		}
		fields := strings.Fields(auth)
		if len(fields) < 2 {
			c.String(http.StatusBadRequest, "no bearer token found\n")
			c.Abort()
			return
		}
		claims, err := utils.ValidateToken(fields[1])
		if err != nil {
			c.String(http.StatusBadRequest, "authentication failed\n")
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.FirstName)
		c.Set("last_name", claims.LastName)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
