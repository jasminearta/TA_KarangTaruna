package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func CreateFotoInovasi(inovasiID uint, fileURL string) (entities.FotoInovasi, error) {
	foto := entities.FotoInovasi{
		InovasiID: inovasiID,
		FileURL:   fileURL,
	}

	err := database.DB.Create(&foto).Error
	return foto, err
}

func GetFotoByInovasi(inovasiID uint) ([]entities.FotoInovasi, error) {
	var foto []entities.FotoInovasi

	err := database.DB.
		Where("inovasi_id = ?", inovasiID).
		Find(&foto).Error

	return foto, err
}
