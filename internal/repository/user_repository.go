package repository

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func CreateUser(user entities.User) (entities.User, error) {
	err := database.DB.Create(&user).Error
	return user, err
}
