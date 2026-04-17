package usecase

import (
	"errors"

	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
	"ta-karangtaruna/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(email, password string) (*entities.User, string, error) {
	var user entities.User

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, "", errors.New("email atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("email atau password salah")
	}

	if user.Status == "nonaktif" {
		return nil, "", errors.New("akun nonaktif, silakan hubungi ketua umum")
	}

	if user.Status == "alumni" {
		return nil, "", errors.New("akun alumni tidak dapat mengakses sistem")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}
