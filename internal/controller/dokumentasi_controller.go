package controllers

import (
	"path/filepath"
	"strconv"

	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

func UploadDokumentasi(c *gin.Context) {

	idParam := c.Param("id")
	kegiatanID, _ := strconv.Atoi(idParam)

	userID := c.MustGet("user_id").(uint)

	// ambil kegiatan
	var kegiatan entities.Kegiatan
	err := database.DB.First(&kegiatan, kegiatanID).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "kegiatan tidak ditemukan",
		})
		return
	}

	// cek apakah kegiatan milik user
	if kegiatan.UserID != userID {
		c.JSON(403, gin.H{
			"error": "anda tidak boleh upload dokumentasi kegiatan orang lain",
		})
		return
	}

	// optional: hanya kegiatan approved
	if kegiatan.Status != "Approved" {
		c.JSON(400, gin.H{
			"error": "dokumentasi hanya bisa diupload setelah kegiatan disetujui",
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

	doc, err := usecases.CreateDokumentasi(uint(kegiatanID), path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// ambil ulang + preload relasi
	var result entities.Dokumentasi

	database.DB.
		Model(&entities.Dokumentasi{}).
		Preload("Kegiatan").
		Preload("Kegiatan.User").
		Preload("Kegiatan.Kategori").
		Where("id = ?", doc.ID).
		First(&result)

	c.JSON(200, gin.H{
		"message": "Dokumentasi berhasil upload",
		"data":    result,
	})
}

func GetDokumentasi(c *gin.Context) {

	idParam := c.Param("id")
	kegiatanID, _ := strconv.Atoi(idParam)

	docs, err := usecases.GetDokumentasiByKegiatan(uint(kegiatanID))

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": docs,
	})
}
