package controllers

import (
	"net/http"
	"strconv"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CreateKegiatanRequest struct {
	Judul      string `json:"judul"`
	Deskripsi  string `json:"deskripsi"`
	Tanggal    string `json:"tanggal"`
	KategoriID uint   `json:"kategori_id"`
}

func CreateKegiatan(c *gin.Context) {

	var req CreateKegiatanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.MustGet("user_id").(uint)

	kegiatan, err := usecases.CreateKegiatan(
		req.Judul,
		req.Deskripsi,
		req.Tanggal,
		req.KategoriID,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Kegiatan berhasil dibuat",
		"data":    kegiatan,
	})
}

func GetKegiatan(c *gin.Context) {

	kegiatan, err := usecases.GetKegiatan()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H

	for _, k := range kegiatan {
		response = append(response, gin.H{
			"ID":         k.ID,
			"UserID":     k.UserID,
			"KategoriID": k.KategoriID,
			"Judul":      k.Judul,
			"Deskripsi":  k.Deskripsi,
			"Tanggal":    k.Tanggal,
			"Status":     k.Status,
			"User":       k.User,
			"Kategori":   k.Kategori,
		})
	}

	c.JSON(200, gin.H{
		"data": response,
	})
}

func GetAllKegiatan(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	search := c.Query("search")
	kategori := c.Query("kategori")
	sortBy := c.Query("sort")
	status := c.Query("status")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * limit

	kegiatan, err := usecases.GetAllKegiatan(limit, offset, search, kategori, sortBy, status)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"page":  page,
		"limit": limit,
		"data":  kegiatan,
	})
}

func GetKegiatanSaya(c *gin.Context) {

	userID := c.MustGet("user_id").(uint)

	kegiatan, err := usecases.GetKegiatanByUser(userID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": kegiatan,
	})
}

func GetAllKegiatanKetua(c *gin.Context) {

	kegiatan, err := usecases.GetKegiatan()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": kegiatan,
	})
}

func GetKegiatanByUser(c *gin.Context) {

	userIDParam := c.Param("id")
	userID, _ := strconv.Atoi(userIDParam)

	kegiatan, err := usecases.GetKegiatanByUser(uint(userID))

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": kegiatan,
	})
}
