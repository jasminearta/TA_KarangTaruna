package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func CreateDokumentasi(kegiatanID uint, fileURL string) (entities.Dokumentasi, error) {

	dokumentasi := entities.Dokumentasi{
		KegiatanID: kegiatanID,
		FileURL:    fileURL,
	}

	err := database.DB.Create(&dokumentasi).Error

	return dokumentasi, err
}
func GetDokumentasiByKegiatan(kegiatanID uint) ([]entities.Dokumentasi, error) {

	var docs []entities.Dokumentasi

	err := database.DB.
		Preload("Kegiatan").
		Preload("Kegiatan.User").
		Preload("Kegiatan.Kategori").
		Where("kegiatan_id = ?", kegiatanID).
		Find(&docs).Error

	return docs, err
}
