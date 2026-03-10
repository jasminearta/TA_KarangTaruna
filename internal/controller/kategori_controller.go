package controllers

import (
	"net/http"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

func GetKategori(c *gin.Context) {

	kategori, err := usecases.GetKategori()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil kategori",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": kategori,
	})
}
