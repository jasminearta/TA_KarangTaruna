package controllers

import (
	"strconv"

	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

func CreateKomentar(c *gin.Context) {

	userID := c.MustGet("user_id").(uint)

	idParam := c.Param("id")
	kegiatanID, _ := strconv.Atoi(idParam)

	var body struct {
		Isi string `json:"isi"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	komentar, err := usecases.CreateKomentar(userID, uint(kegiatanID), body.Isi)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// ambil ulang komentar + relasi
	var result entities.Komentar

	database.DB.
		Preload("User").
		Preload("Kegiatan").
		First(&result, komentar.ID)

	c.JSON(200, gin.H{
		"message": "Komentar berhasil dibuat",
		"data":    result,
	})
}
func GetKomentar(c *gin.Context) {

	idParam := c.Param("id")
	kegiatanID, _ := strconv.Atoi(idParam)

	komentar, err := usecases.GetKomentarByKegiatan(uint(kegiatanID))

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": komentar,
	})
}
