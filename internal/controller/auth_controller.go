package controllers

import (
	"fmt"
	"net/http"

	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegisterRequest struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Register ketua divisi
// @Description Mendaftarkan user dengan role ketua_divisi
// @Tags Auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body entities.User true "Data register"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
	var input RegisterRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("REGISTER NAMA:", input.Nama)
	fmt.Println("REGISTER EMAIL:", input.Email)
	fmt.Println("REGISTER PASSWORD:", input.Password)

	newUser, err := usecases.RegisterUser(
		input.Nama,
		input.Email,
		input.Password,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User berhasil register",
		"data": gin.H{
			"id":     newUser.ID,
			"nama":   newUser.Nama,
			"email":  newUser.Email,
			"role":   newUser.Role,
			"status": newUser.Status,
		},
	})
}

// @Summary Register ketua umum
// @Description Mendaftarkan user dengan role ketua_umum
// @Tags Auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body entities.User true "Data register ketua umum"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /register-ketua [post]
func RegisterKetua(c *gin.Context) {
	var input RegisterRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("REGISTER KETUA NAMA:", input.Nama)
	fmt.Println("REGISTER KETUA EMAIL:", input.Email)
	fmt.Println("REGISTER KETUA PASSWORD:", input.Password)

	newUser, err := usecases.RegisterKetua(
		input.Nama,
		input.Email,
		input.Password,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ketua berhasil dibuat",
		"data": gin.H{
			"id":     newUser.ID,
			"nama":   newUser.Nama,
			"email":  newUser.Email,
			"role":   newUser.Role,
			"status": newUser.Status,
		},
	})
}

// @Summary Login user
// @Description Login untuk ketua divisi atau ketua umum
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body controllers.LoginRequest true "Login payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	var input LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, token, err := usecases.LoginUser(input.Email, input.Password)
	if err != nil {
		if err.Error() == "akun nonaktif, silakan hubungi ketua umum" ||
			err.Error() == "akun alumni tidak dapat mengakses sistem" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"user": gin.H{
			"id":     user.ID,
			"nama":   user.Nama,
			"email":  user.Email,
			"foto":   user.Foto,
			"role":   user.Role,
			"status": user.Status,
		},
		"token": token,
	})
}
