package usecase

import (
	"errors"
	"strings"

	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func UpdateUserStatus(userID uint, status string) (*entities.User, error) {
	status = strings.ToLower(strings.TrimSpace(status))

	if status != "aktif" && status != "nonaktif" && status != "alumni" {
		return nil, errors.New("status user tidak valid")
	}

	var user entities.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	user.Status = status

	if err := database.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
