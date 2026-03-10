package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func GetUserByID(userID uint) (entities.User, error) {

	var user entities.User

	err := database.DB.First(&user, userID).Error

	return user, err
}

func GetAllUsers() ([]entities.User, error) {

	var users []entities.User

	err := database.DB.Find(&users).Error

	return users, err
}
