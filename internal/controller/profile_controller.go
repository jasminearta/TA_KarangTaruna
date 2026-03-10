package controllers

import (
	"net/http"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {

	userID := c.MustGet("user_id").(uint)

	user, err := usecases.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"nama":  user.Nama,
		"email": user.Email,
		"role":  user.Role,
	})
}
