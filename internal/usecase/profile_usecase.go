package usecase

import (
	"errors"

	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"

	"golang.org/x/crypto/bcrypt"
)

func UpdateProfile(userID uint, nama string, email string) (entities.User, error) {
	var user entities.User

	err := database.DB.First(&user, userID).Error
	if err != nil {
		return user, errors.New("user tidak ditemukan")
	}

	user.Nama = nama
	user.Email = email

	err = database.DB.Save(&user).Error
	return user, err
}

func ChangePassword(userID uint, oldPassword string, newPassword string) error {
	var user entities.User

	err := database.DB.First(&user, userID).Error
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return errors.New("password lama salah")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = database.DB.Save(&user).Error
	return err
}

func UpdateFotoProfile(userID uint, foto string) (entities.User, error) {
	var user entities.User

	err := database.DB.First(&user, userID).Error
	if err != nil {
		return user, errors.New("user tidak ditemukan")
	}

	user.Foto = foto

	err = database.DB.Save(&user).Error
	return user, err
}
