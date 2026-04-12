package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func CreateFotoKegiatan(kegiatanID uint, fileURL string) (entities.FotoKegiatan, error) {
	foto := entities.FotoKegiatan{
		KegiatanID: kegiatanID,
		FileURL:    fileURL,
	}

	err := database.DB.Create(&foto).Error
	return foto, err
}

func GetFotoByKegiatan(kegiatanID uint) ([]entities.FotoKegiatan, error) {
	var foto []entities.FotoKegiatan

	err := database.DB.
		Where("kegiatan_id = ?", kegiatanID).
		Find(&foto).Error

	return foto, err
}
