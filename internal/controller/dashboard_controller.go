package controllers

import (
	"net/http"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

func GetDashboardKetua(c *gin.Context) {

	data, err := usecases.GetDashboardData()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
