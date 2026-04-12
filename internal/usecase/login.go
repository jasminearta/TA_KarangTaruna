package usecase

import (
	"errors"
	"fmt"

	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
	"ta-karangtaruna/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(email string, password string) (entities.User, string, error) {
	var user entities.User

	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		fmt.Println("USER TIDAK DITEMUKAN")
		return user, "", errors.New("email tidak ditemukan")
	}

	fmt.Println("EMAIL DB:", user.Email)
	fmt.Println("PASSWORD HASH DB:", user.Password)
	fmt.Println("PASSWORD INPUT:", password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("PASSWORD TIDAK MATCH")
		return user, "", errors.New("password salah")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return user, "", err
	}

	return user, token, nil
}
