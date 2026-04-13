package controllers

import (
	"net/http"
	"strconv"
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"

	"github.com/gin-gonic/gin"
)

// @Summary Get notifications
// @Description Mengambil semua notifikasi milik user yang sedang login
// @Tags Notifications
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notifications [get]
func GetNotifications(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var notifications []entities.Notification

	err := database.DB.
		Preload("User").
		Where("user_id = ?", userID).
		Order("id desc").
		Find(&notifications).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": notifications,
	})
}

// @Summary Read notification
// @Description Menandai notifikasi sebagai sudah dibaca
// @Tags Notifications
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID Notification"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notifications/{id}/read [patch]
func ReadNotification(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id notification tidak valid",
		})
		return
	}

	userID := c.MustGet("user_id").(uint)

	err = database.DB.
		Model(&entities.Notification{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("is_read", true).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Notification read",
	})
}
