package controllers

import (
	"net/http"
	"strconv"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UpdateUserStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

// @Summary Update status user
// @Description Ketua umum mengubah status user menjadi aktif, nonaktif, atau alumni
// @Tags Ketua Umum User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID User"
// @Param request body controllers.UpdateUserStatusRequest true "Data status user"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ketua/users/{id}/status [patch]
func UpdateUserStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID user tidak valid",
		})
		return
	}

	var req UpdateUserStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := usecases.UpdateUserStatus(uint(id), req.Status)
	if err != nil {
		if err.Error() == "status user tidak valid" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "user tidak ditemukan" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Status user berhasil diperbarui",
		"data": gin.H{
			"id":     user.ID,
			"nama":   user.Nama,
			"email":  user.Email,
			"role":   user.Role,
			"status": user.Status,
		},
	})
}
