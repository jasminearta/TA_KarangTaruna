package controllers

import (
	"net/http"
	"strconv"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CreateInovasiRequest struct {
	Judul           string `json:"judul"`
	Deskripsi       string `json:"deskripsi"`
	TanggalDiajukan string `json:"tanggal_diajukan"`
	KategoriID      uint   `json:"kategori_id"`
}

func CreateInovasi(c *gin.Context) {
	var req CreateInovasiRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.MustGet("user_id").(uint)

	inovasi, err := usecases.CreateInovasi(
		req.Judul,
		req.Deskripsi,
		req.TanggalDiajukan,
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
		"message": "Inovasi berhasil dibuat",
		"data":    inovasi,
	})
}

func GetInovasi(c *gin.Context) {
	inovasi, err := usecases.GetInovasi()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H

	for _, i := range inovasi {
		response = append(response, gin.H{
			"ID":              i.ID,
			"UserID":          i.UserID,
			"KategoriID":      i.KategoriID,
			"Judul":           i.Judul,
			"Deskripsi":       i.Deskripsi,
			"TanggalDiajukan": i.TanggalDiajukan,
			"Status":          i.Status,
			"User":            i.User,
			"Kategori":        i.Kategori,
		})
	}

	c.JSON(200, gin.H{
		"data": response,
	})
}

func GetAllInovasi(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "100")

	search := c.Query("search")
	kategori := c.Query("kategori")
	sortBy := c.Query("sort")
	status := c.Query("status")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * limit

	inovasi, err := usecases.GetAllInovasi(limit, offset, search, kategori, sortBy, status)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"page":  page,
		"limit": limit,
		"data":  inovasi,
	})
}

func GetInovasiSaya(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "100")

	search := c.Query("search")
	kategori := c.Query("kategori")
	sortBy := c.Query("sort")
	status := c.Query("status")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	offset := (page - 1) * limit

	inovasi, err := usecases.GetInovasiSayaUsecase(userID, limit, offset, search, kategori, sortBy, status)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H

	for _, i := range inovasi {
		item := gin.H{
			"id":               i.ID,
			"user_id":          i.UserID,
			"kategori_id":      i.KategoriID,
			"judul":            i.Judul,
			"deskripsi":        i.Deskripsi,
			"tanggal_diajukan": i.TanggalDiajukan,
			"status":           i.Status,
			"user":             i.User,
			"kategori":         i.Kategori,
			"foto_inovasi":     i.FotoInovasi,
		}

		if i.Status == "rejected" {
			log, err := usecases.GetLatestApprovalLogByTarget(i.ID, "inovasi")
			if err == nil {
				item["catatan_penolakan"] = log.Catatan
			}
		}

		response = append(response, item)
	}

	c.JSON(200, gin.H{
		"page":  page,
		"limit": limit,
		"data":  response,
	})
}

func GetAllInovasiKetua(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "100")

	search := c.Query("search")
	kategori := c.Query("kategori")
	sortBy := c.Query("sort")
	status := c.Query("status")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * limit

	inovasi, err := usecases.GetAllInovasiKetuaUsecase(limit, offset, search, kategori, sortBy, status)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H

	for _, i := range inovasi {
		item := gin.H{
			"id":               i.ID,
			"user_id":          i.UserID,
			"kategori_id":      i.KategoriID,
			"judul":            i.Judul,
			"deskripsi":        i.Deskripsi,
			"tanggal_diajukan": i.TanggalDiajukan,
			"status":           i.Status,
			"user":             i.User,
			"kategori":         i.Kategori,
			"foto_inovasi":     i.FotoInovasi,
		}

		if i.Status == "rejected" {
			log, err := usecases.GetLatestApprovalLogByTarget(i.ID, "inovasi")
			if err == nil {
				item["catatan_penolakan"] = log.Catatan
			}
		}

		response = append(response, item)
	}

	c.JSON(200, gin.H{
		"page":  page,
		"limit": limit,
		"data":  response,
	})
}

func GetInovasiByUser(c *gin.Context) {
	userIDParam := c.Param("id")
	userID, _ := strconv.Atoi(userIDParam)

	inovasi, err := usecases.GetInovasiByUser(uint(userID))

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": inovasi,
	})
}

func GetDetailInovasi(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	inovasi, err := usecases.GetDetailInovasi(uint(id))

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": inovasi,
	})
}

func UpdateInovasi(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var req CreateInovasiRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uint)

	inovasi, err := usecases.UpdateInovasi(
		uint(id),
		req.Judul,
		req.Deskripsi,
		req.TanggalDiajukan,
		req.KategoriID,
		userID,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Inovasi berhasil diupdate",
		"data":    inovasi,
	})
}

func DeleteInovasi(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	userID := c.MustGet("user_id").(uint)

	err := usecases.DeleteInovasi(uint(id), userID)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Inovasi berhasil dihapus",
	})
}
