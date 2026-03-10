package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func CreateKomentar(userID uint, kegiatanID uint, isi string) (entities.Komentar, error) {

	komentar := entities.Komentar{
		UserID:     userID,
		KegiatanID: kegiatanID,
		Isi:        isi,
	}

	err := database.DB.Create(&komentar).Error

	return komentar, err
}

func GetKomentarByKegiatan(kegiatanID uint) ([]entities.Komentar, error) {

	var komentar []entities.Komentar

	err := database.DB.
		Preload("User").
		Preload("Kegiatan").
		Preload("Kegiatan.Kategori").
		Where("kegiatan_id = ?", kegiatanID).
		Order("id asc").
		Find(&komentar).Error

	return komentar, err
}
