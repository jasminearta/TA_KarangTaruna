package controllers

import (
	"net/http"
	"path/filepath"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UpdateProfileRequest struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

type ChangePasswordRequest struct {
	PasswordLama string `json:"password_lama"`
	PasswordBaru string `json:"password_baru"`
}

// @Summary Get profile
// @Description Mengambil profile user yang sedang login
// @Tags Profile
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/profile [get]
func GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	user, err := usecases.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "User tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"nama":  user.Nama,
		"email": user.Email,
		"foto":  user.Foto,
		"role":  user.Role,
	})
}

// @Summary Update profile
// @Description Mengubah nama dan email user yang sedang login
// @Tags Profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body controllers.UpdateProfileRequest true "Data update profile"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/profile [put]
func UpdateProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := usecases.UpdateProfile(userID, req.Nama, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile berhasil diupdate",
		"data": gin.H{
			"id":    user.ID,
			"nama":  user.Nama,
			"email": user.Email,
			"foto":  user.Foto,
			"role":  user.Role,
		},
	})
}

// @Summary Change password
// @Description Mengubah password user yang sedang login
// @Tags Profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body controllers.ChangePasswordRequest true "Data ubah password"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/profile/password [put]
func ChangePassword(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := usecases.ChangePassword(userID, req.PasswordLama, req.PasswordBaru)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password berhasil diubah",
	})
}

// @Summary Upload foto profile
// @Description Upload foto profile user yang sedang login
// @Tags Profile
// @Security BearerAuth
// @Accept mpfd
// @Produce json
// @Param file formData file true "File foto profile"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/profile/foto [post]
func UploadFotoProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

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

	user, err := usecases.UpdateFotoProfile(userID, path)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Foto profile berhasil diupload",
		"data": gin.H{
			"id":    user.ID,
			"nama":  user.Nama,
			"email": user.Email,
			"foto":  user.Foto,
			"role":  user.Role,
		},
	})
}
