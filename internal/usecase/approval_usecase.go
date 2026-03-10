package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func UpdateStatusKegiatan(id uint, status string) error {

	var kegiatan entities.Kegiatan

	err := database.DB.First(&kegiatan, id).Error
	if err != nil {
		return err
	}

	kegiatan.Status = status

	return database.DB.Save(&kegiatan).Error
}
