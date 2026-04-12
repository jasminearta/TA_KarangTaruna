package controllers

import (
	"net/http"
	"strconv"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CreateKegiatanRequest struct {
	Judul           string `json:"judul"`
	Deskripsi       string `json:"deskripsi"`
	TanggalBerjalan string `json:"tanggal_berjalan"`
	TanggalDiajukan string `json:"tanggal_diajukan"`
	KategoriID      uint   `json:"kategori_id"`
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
		req.TanggalBerjalan,
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
		item := gin.H{
			"id":               k.ID,
			"user_id":          k.UserID,
			"kategori_id":      k.KategoriID,
			"judul":            k.Judul,
			"deskripsi":        k.Deskripsi,
			"tanggal_berjalan": k.TanggalBerjalan,
			"tanggal_diajukan": k.TanggalDiajukan,
			"status":           k.Status,
			"user":             k.User,
			"kategori":         k.Kategori,
		}

		response = append(response, item)
	}

	c.JSON(200, gin.H{
		"data": response,
	})
}

func GetAllKegiatan(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "100")

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

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "100")

	search := c.Query("search")
	kategori := c.Query("kategori")
	sortBy := c.Query("sort")
	status := c.Query("status")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	offset := (page - 1) * limit

	kegiatan, err := usecases.GetKegiatanSayaUsecase(userID, limit, offset, search, kategori, sortBy, status)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H

	for _, k := range kegiatan {
		item := gin.H{
			"id":               k.ID,
			"user_id":          k.UserID,
			"kategori_id":      k.KategoriID,
			"judul":            k.Judul,
			"deskripsi":        k.Deskripsi,
			"tanggal_berjalan": k.TanggalBerjalan,
			"tanggal_diajukan": k.TanggalDiajukan,
			"status":           k.Status,
			"user":             k.User,
			"kategori":         k.Kategori,
			"foto_kegiatan":    k.FotoKegiatan,
		}

		if k.Status == "rejected" {
			log, err := usecases.GetLatestApprovalLogByTarget(k.ID, "kegiatan")
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

func GetAllKegiatanKetua(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "100")

	search := c.Query("search")
	kategori := c.Query("kategori")
	sortBy := c.Query("sort")
	status := c.Query("status")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	offset := (page - 1) * limit

	kegiatan, err := usecases.GetAllKegiatanKetuaUsecase(limit, offset, search, kategori, sortBy, status)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H

	for _, k := range kegiatan {
		item := gin.H{
			"id":               k.ID,
			"user_id":          k.UserID,
			"kategori_id":      k.KategoriID,
			"judul":            k.Judul,
			"deskripsi":        k.Deskripsi,
			"tanggal_berjalan": k.TanggalBerjalan,
			"tanggal_diajukan": k.TanggalDiajukan,
			"status":           k.Status,
			"user":             k.User,
			"kategori":         k.Kategori,
			"foto_kegiatan":    k.FotoKegiatan,
		}

		if k.Status == "rejected" {
			log, err := usecases.GetLatestApprovalLogByTarget(k.ID, "kegiatan")
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

func GetKegiatanByUser(c *gin.Context) {
	userIDParam := c.Param("id")
	userID, _ := strconv.Atoi(userIDParam)

	kegiatan, err := usecases.GetKegiatanByUser(uint(userID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H

	for _, k := range kegiatan {
		item := gin.H{
			"id":               k.ID,
			"user_id":          k.UserID,
			"kategori_id":      k.KategoriID,
			"judul":            k.Judul,
			"deskripsi":        k.Deskripsi,
			"tanggal_berjalan": k.TanggalBerjalan,
			"tanggal_diajukan": k.TanggalDiajukan,
			"status":           k.Status,
			"user":             k.User,
			"kategori":         k.Kategori,
			"foto_kegiatan":    k.FotoKegiatan,
		}

		if k.Status == "rejected" {
			log, err := usecases.GetLatestApprovalLogByTarget(k.ID, "kegiatan")
			if err == nil {
				item["catatan_penolakan"] = log.Catatan
			}
		}

		response = append(response, item)
	}

	c.JSON(200, gin.H{
		"data": response,
	})
}

func GetDetailKegiatan(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	kegiatan, err := usecases.GetDetailKegiatan(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	item := gin.H{
		"id":               kegiatan.ID,
		"user_id":          kegiatan.UserID,
		"kategori_id":      kegiatan.KategoriID,
		"judul":            kegiatan.Judul,
		"deskripsi":        kegiatan.Deskripsi,
		"tanggal_berjalan": kegiatan.TanggalBerjalan,
		"tanggal_diajukan": kegiatan.TanggalDiajukan,
		"status":           kegiatan.Status,
		"user":             kegiatan.User,
		"kategori":         kegiatan.Kategori,
		"foto_kegiatan":    kegiatan.FotoKegiatan,
	}

	if kegiatan.Status == "rejected" {
		log, err := usecases.GetLatestApprovalLogByTarget(kegiatan.ID, "kegiatan")
		if err == nil {
			item["catatan_penolakan"] = log.Catatan
		}
	}

	c.JSON(200, gin.H{
		"data": item,
	})
}

func UpdateKegiatan(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var req CreateKegiatanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uint)

	kegiatan, err := usecases.UpdateKegiatan(
		uint(id),
		req.Judul,
		req.Deskripsi,
		req.TanggalBerjalan,
		req.TanggalDiajukan,
		req.KategoriID,
		userID,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Kegiatan berhasil diupdate",
		"data":    kegiatan,
	})
}

func DeleteKegiatan(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	userID := c.MustGet("user_id").(uint)

	err := usecases.DeleteKegiatan(uint(id), userID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Kegiatan berhasil dihapus",
	})
}
