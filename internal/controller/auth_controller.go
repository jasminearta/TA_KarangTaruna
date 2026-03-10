package controllers

import (
	"net/http"

	"ta-karangtaruna/internal/entities"
	usecases "ta-karangtaruna/internal/usecase"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var input entities.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
			"id":    newUser.ID,
			"nama":  newUser.Nama,
			"email": newUser.Email,
			"role":  newUser.Role,
		},
	})
}

func RegisterKetua(c *gin.Context) {

	var input entities.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
			"id":    newUser.ID,
			"nama":  newUser.Nama,
			"email": newUser.Email,
			"role":  newUser.Role,
		},
	})
}

func Login(c *gin.Context) {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, token, err := usecases.LoginUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"user_id": user.ID,
		"role":    user.Role,
		"token":   token,
	})
}
