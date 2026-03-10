package usecase

import (
	"time"

	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func CreateKegiatan(
	judul string,
	deskripsi string,
	tanggal string,
	kategoriID uint,
	userID uint,
) (entities.Kegiatan, error) {

	parsedDate, _ := time.Parse("2006-01-02", tanggal)

	kegiatan := entities.Kegiatan{
		Judul:      judul,
		Deskripsi:  deskripsi,
		Tanggal:    parsedDate,
		KategoriID: kategoriID,
		UserID:     userID,
		Status:     "pending",
	}

	err := database.DB.Create(&kegiatan).Error
	if err != nil {
		return kegiatan, err
	}

	database.DB.
		Preload("User").
		Preload("Kategori").
		First(&kegiatan, kegiatan.ID)

	return kegiatan, nil
}
