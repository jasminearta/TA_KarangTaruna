package controllers

import (
	"net/http"
	"strconv"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

func ApproveKegiatan(c *gin.Context) {

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := usecases.UpdateStatusKegiatan(uint(id), "approved")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Kegiatan berhasil di-approve",
	})
}

func RejectKegiatan(c *gin.Context) {

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := usecases.UpdateStatusKegiatan(uint(id), "rejected")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Kegiatan berhasil di-reject",
	})
}
