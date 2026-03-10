package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OnlyKetua() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")

		if role != "ketua" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Hanya ketua yang boleh akses",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
