package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func CreateKegiatan(
	judul string,
	deskripsi string,
	tanggalBerjalan string,
	tanggalDiajukan string,
	kategoriID uint,
	userID uint,
) (entities.Kegiatan, error) {

	kegiatan := entities.Kegiatan{
		Judul:           judul,
		Deskripsi:       deskripsi,
		TanggalBerjalan: tanggalBerjalan,
		TanggalDiajukan: tanggalDiajukan,
		KategoriID:      kategoriID,
		UserID:          userID,
		Status:          "pending",
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
