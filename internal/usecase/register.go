package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(nama string, email string, password string) (entities.User, error) {

	user := entities.User{
		Nama:  nama,
		Email: email,
		Role:  "anggota", // ✅ otomatis
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Password = string(hashedPassword)

	err = database.DB.Create(&user).Error
	return user, err
}
func RegisterKetua(nama string, email string, password string) (entities.User, error) {

	user := entities.User{
		Nama:  nama,
		Email: email,
		Role:  "ketua",
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Password = string(hashedPassword)

	err = database.DB.Create(&user).Error

	return user, err
}
