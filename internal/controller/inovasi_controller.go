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

// @Summary Create inovasi
// @Description Ketua divisi membuat pengajuan inovasi
// @Tags Ketua Divisi Inovasi
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body controllers.CreateInovasiRequest true "Data create inovasi"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/inovasi [post]
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

// @Summary Get all public inovasi
// @Description Mengambil daftar inovasi publik yang berstatus approved
// @Tags Public Inovasi
// @Produce json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param search query string false "Search judul"
// @Param kategori query string false "Kategori ID"
// @Param sort query string false "Sort: terbaru atau terlama"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /inovasi [get]
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

// @Summary Get inovasi saya
// @Description Mengambil daftar inovasi milik ketua divisi
// @Tags Ketua Divisi Inovasi
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param search query string false "Search judul"
// @Param kategori query string false "Kategori ID"
// @Param status query string false "Status: pending, approved, rejected"
// @Param sort query string false "Sort: terbaru atau terlama"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/inovasi-saya [get]
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

// @Summary Get all inovasi for ketua umum
// @Description Mengambil semua inovasi untuk ketua umum
// @Tags Ketua Umum Inovasi
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param search query string false "Search judul"
// @Param kategori query string false "Kategori ID"
// @Param status query string false "Status: pending, approved, rejected"
// @Param sort query string false "Sort: terbaru atau terlama"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ketua/inovasi [get]
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

// @Summary Get inovasi by user
// @Description Mengambil inovasi berdasarkan user tertentu
// @Tags Ketua Umum Inovasi
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID User"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ketua/inovasi/user/{id} [get]
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

// @Summary Get detail inovasi
// @Description Mengambil detail inovasi publik berdasarkan ID
// @Tags Public Inovasi
// @Produce json
// @Param id path int true "ID Inovasi"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /inovasi/{id} [get]
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

// @Summary Update inovasi
// @Description Mengubah data inovasi milik ketua divisi selama status masih pending
// @Tags Ketua Divisi Inovasi
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID Inovasi"
// @Param request body controllers.CreateInovasiRequest true "Data update inovasi"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/inovasi/{id} [put]
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

// @Summary Delete inovasi
// @Description Menghapus inovasi milik ketua divisi selama status masih pending
// @Tags Ketua Divisi Inovasi
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID Inovasi"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/inovasi/{id} [delete]
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
