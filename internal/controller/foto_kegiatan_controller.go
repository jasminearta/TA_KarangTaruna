package controllers

import (
	"path/filepath"
	"strconv"

	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

// @Summary Upload foto kegiatan
// @Description Upload foto kegiatan setelah kegiatan disetujui
// @Tags Ketua Divisi Kegiatan
// @Security BearerAuth
// @Accept mpfd
// @Produce json
// @Param id path int true "ID Kegiatan"
// @Param file formData file true "File foto kegiatan"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/kegiatan/{id}/foto [post]
func UploadFotoKegiatan(c *gin.Context) {
	idParam := c.Param("id")
	kegiatanID, _ := strconv.Atoi(idParam)

	userID := c.MustGet("user_id").(uint)

	var kegiatan entities.Kegiatan
	err := database.DB.First(&kegiatan, kegiatanID).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "kegiatan tidak ditemukan",
		})
		return
	}

	if kegiatan.UserID != userID {
		c.JSON(403, gin.H{
			"error": "anda tidak boleh upload foto kegiatan orang lain",
		})
		return
	}

	if kegiatan.Status != "approved" {
		c.JSON(400, gin.H{
			"error": "foto kegiatan hanya bisa diupload setelah kegiatan disetujui",
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

	foto, err := usecases.CreateFotoKegiatan(uint(kegiatanID), path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var result entities.FotoKegiatan
	database.DB.
		Model(&entities.FotoKegiatan{}).
		Preload("Kegiatan").
		Preload("Kegiatan.User").
		Preload("Kegiatan.Kategori").
		Where("id = ?", foto.ID).
		First(&result)

	c.JSON(200, gin.H{
		"message": "Foto kegiatan berhasil upload",
		"data":    result,
	})
}

// @Summary Get foto kegiatan
// @Description Mengambil daftar foto kegiatan berdasarkan ID kegiatan
// @Tags Public Kegiatan
// @Produce json
// @Param id path int true "ID Kegiatan"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /kegiatan/{id}/foto [get]
func GetFotoKegiatan(c *gin.Context) {
	idParam := c.Param("id")
	kegiatanID, _ := strconv.Atoi(idParam)

	foto, err := usecases.GetFotoByKegiatan(uint(kegiatanID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": foto,
	})
}
