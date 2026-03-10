package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func GetKegiatan() ([]entities.Kegiatan, error) {

	var kegiatan []entities.Kegiatan

	err := database.DB.
		Preload("User").
		Preload("Kategori").
		Find(&kegiatan).Error

	return kegiatan, err
}

func GetAllKegiatan(limit int, offset int, search string, kategori string, sort string, status string) ([]entities.Kegiatan, error) {

	var kegiatan []entities.Kegiatan

	query := database.DB.
		Preload("User").
		Preload("Kategori")

	if search != "" {
		query = query.Where("judul LIKE ?", "%"+search+"%")
	}

	if kategori != "" {
		query = query.Where("kategori_id = ?", kategori)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if sort == "terbaru" {
		query = query.Order("tanggal DESC")
	}

	if sort == "terlama" {
		query = query.Order("tanggal ASC")
	}

	err := query.
		Limit(limit).
		Offset(offset).
		Find(&kegiatan).Error

	return kegiatan, err
}

func GetKegiatanByUser(userID uint) ([]entities.Kegiatan, error) {

	var kegiatan []entities.Kegiatan

	err := database.DB.
		Preload("User").
		Preload("Kategori").
		Where("user_id = ?", userID).
		Find(&kegiatan).Error

	return kegiatan, err
}
