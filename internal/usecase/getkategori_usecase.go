package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func GetKategori() ([]entities.Kategori, error) {

	var kategori []entities.Kategori

	err := database.DB.Find(&kategori).Error

	return kategori, err
}
