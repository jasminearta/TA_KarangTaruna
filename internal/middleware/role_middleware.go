package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OnlyKetuaUmum() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")

		if role != "ketua_umum" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Hanya ketua umum yang boleh akses",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func OnlyKetuaDivisi() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")

		if role != "ketua_divisi" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Hanya ketua divisi yang boleh akses",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
