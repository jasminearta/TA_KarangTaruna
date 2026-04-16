package controllers

import (
	"net/http"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

// @Summary Get dashboard ketua umum
// @Description Mengambil data dashboard ketua umum berupa total kegiatan, total inovasi, total pengguna, serta jumlah data berdasarkan status
// @Tags Ketua Umum Dashboard
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ketua/dashboard [get]
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
