package controllers

import (
	"net/http"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {

	users, err := usecases.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
