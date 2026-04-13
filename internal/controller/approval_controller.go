package controllers

import (
	"net/http"
	"strconv"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RejectRequest struct {
	Catatan string `json:"catatan"`
}

// @Summary Approve kegiatan
// @Description Menyetujui pengajuan kegiatan
// @Tags Ketua Umum Kegiatan
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID Kegiatan"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ketua/kegiatan/{id}/approve [patch]
func ApproveKegiatan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id tidak valid",
		})
		return
	}

	approvedBy := c.MustGet("user_id").(uint)

	kegiatan, err := usecases.ApproveKegiatanUsecase(uint(id), approvedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = usecases.CreateNotification(
		kegiatan.UserID,
		"Kegiatan Disetujui",
		"Kegiatan '"+kegiatan.Judul+"' telah disetujui oleh ketua",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "status berubah, tapi notifikasi gagal dibuat: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Kegiatan berhasil di-approve",
	})
}

// @Summary Reject kegiatan
// @Description Menolak pengajuan kegiatan disertai catatan
// @Tags Ketua Umum Kegiatan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID Kegiatan"
// @Param request body controllers.RejectRequest true "Catatan penolakan"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ketua/kegiatan/{id}/reject [patch]
func RejectKegiatan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id tidak valid",
		})
		return
	}

	var req RejectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	approvedBy := c.MustGet("user_id").(uint)

	kegiatan, err := usecases.RejectKegiatanUsecase(uint(id), approvedBy, req.Catatan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = usecases.CreateNotification(
		kegiatan.UserID,
		"Kegiatan Ditolak",
		"Kegiatan '"+kegiatan.Judul+"' ditolak oleh ketua",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "status berubah, tapi notifikasi gagal dibuat: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Kegiatan berhasil di-reject",
	})
}

func ApproveInovasi(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id tidak valid",
		})
		return
	}

	approvedBy := c.MustGet("user_id").(uint)

	inovasi, err := usecases.ApproveInovasiUsecase(uint(id), approvedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = usecases.CreateNotification(
		inovasi.UserID,
		"Inovasi Disetujui",
		"Inovasi '"+inovasi.Judul+"' telah disetujui oleh ketua",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "status berubah, tapi notifikasi gagal dibuat: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Inovasi berhasil di-approve",
	})
}

func RejectInovasi(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id tidak valid",
		})
		return
	}

	var req RejectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	approvedBy := c.MustGet("user_id").(uint)

	inovasi, err := usecases.RejectInovasiUsecase(uint(id), approvedBy, req.Catatan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = usecases.CreateNotification(
		inovasi.UserID,
		"Inovasi Ditolak",
		"Inovasi '"+inovasi.Judul+"' ditolak oleh ketua",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "status berubah, tapi notifikasi gagal dibuat: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Inovasi berhasil di-reject",
	})
}
