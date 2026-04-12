package controllers

import (
	"path/filepath"
	"strconv"

	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

func UploadFotoInovasi(c *gin.Context) {
	idParam := c.Param("id")
	inovasiID, _ := strconv.Atoi(idParam)

	userID := c.MustGet("user_id").(uint)

	var inovasi entities.Inovasi
	err := database.DB.First(&inovasi, inovasiID).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "inovasi tidak ditemukan",
		})
		return
	}

	if inovasi.UserID != userID {
		c.JSON(403, gin.H{
			"error": "anda tidak boleh upload foto inovasi orang lain",
		})
		return
	}

	if inovasi.Status != "approved" {
		c.JSON(400, gin.H{
			"error": "foto inovasi hanya bisa diupload setelah inovasi disetujui",
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "file tidak ditemukan",
		})
		return
	}

	filename := filepath.Base(file.Filename)
	path := "uploads/" + filename

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "gagal upload file",
		})
		return
	}

	foto, err := usecases.CreateFotoInovasi(uint(inovasiID), path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var result entities.FotoInovasi
	database.DB.
		Model(&entities.FotoInovasi{}).
		Where("id = ?", foto.ID).
		First(&result)

	c.JSON(200, gin.H{
		"message": "Foto inovasi berhasil upload",
		"data":    result,
	})
}

func GetFotoInovasi(c *gin.Context) {
	idParam := c.Param("id")
	inovasiID, _ := strconv.Atoi(idParam)

	foto, err := usecases.GetFotoByInovasi(uint(inovasiID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": foto,
	})
}
