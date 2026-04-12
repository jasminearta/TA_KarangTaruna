package usecase

import (
	"errors"
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func GetKegiatan() ([]entities.Kegiatan, error) {
	var kegiatan []entities.Kegiatan

	err := database.DB.
		Preload("User").
		Preload("Kategori").
		Where("status = ?", "approved").
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
		query = query.Where("LOWER(status) = ?", status)
	}

	if sort == "pending" || sort == "approved" || sort == "rejected" {
		query = query.Where("LOWER(status) = ?", sort)
	}

	if status == "" && sort != "pending" && sort != "approved" && sort != "rejected" {
		query = query.Where("LOWER(status) = ?", "approved")
	}

	if sort == "terbaru" {
		query = query.Order("tanggal_diajukan DESC")
	}

	if sort == "terlama" {
		query = query.Order("tanggal_diajukan ASC")
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
		Preload("FotoKegiatan").
		Where("user_id = ?", userID).
		Find(&kegiatan).Error

	return kegiatan, err
}

func GetDetailKegiatan(id uint) (entities.Kegiatan, error) {
	var kegiatan entities.Kegiatan

	err := database.DB.
		Preload("User").
		Preload("Kategori").
		Preload("FotoKegiatan").
		Where("status = ?", "approved").
		First(&kegiatan, id).Error

	return kegiatan, err
}

func UpdateKegiatan(
	id uint,
	judul string,
	deskripsi string,
	tanggalBerjalan string,
	tanggalDiajukan string,
	kategoriID uint,
	userID uint,
) (entities.Kegiatan, error) {

	var kegiatan entities.Kegiatan

	err := database.DB.First(&kegiatan, id).Error
	if err != nil {
		return kegiatan, err
	}

	if kegiatan.UserID != userID {
		return kegiatan, errors.New("anda tidak memiliki akses untuk mengedit kegiatan ini")
	}

	if kegiatan.Status != "pending" {
		return kegiatan, errors.New("kegiatan yang sudah diverifikasi tidak dapat diubah")
	}

	kegiatan.Judul = judul
	kegiatan.Deskripsi = deskripsi
	kegiatan.TanggalBerjalan = tanggalBerjalan
	kegiatan.TanggalDiajukan = tanggalDiajukan
	kegiatan.KategoriID = kategoriID

	err = database.DB.Save(&kegiatan).Error
	if err != nil {
		return kegiatan, err
	}

	err = database.DB.
		Preload("User").
		Preload("Kategori").
		Preload("FotoKegiatan").
		First(&kegiatan, kegiatan.ID).Error

	return kegiatan, err
}

func DeleteKegiatan(id uint, userID uint) error {
	var kegiatan entities.Kegiatan

	err := database.DB.First(&kegiatan, id).Error
	if err != nil {
		return errors.New("kegiatan tidak ditemukan")
	}

	if kegiatan.UserID != userID {
		return errors.New("anda tidak memiliki akses")
	}

	if kegiatan.Status != "pending" {
		return errors.New("kegiatan yang sudah diverifikasi tidak bisa dihapus")
	}

	if err := database.DB.Where("kegiatan_id = ?", id).Delete(&entities.FotoKegiatan{}).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&kegiatan).Error; err != nil {
		return err
	}

	return nil
}

func GetAllKegiatanKetuaUsecase(limit int, offset int, search string, kategori string, sort string, status string) ([]entities.Kegiatan, error) {
	var kegiatan []entities.Kegiatan

	query := database.DB.
		Preload("User").
		Preload("Kategori").
		Preload("FotoKegiatan")

	if search != "" {
		query = query.Where("judul LIKE ?", "%"+search+"%")
	}

	if kategori != "" {
		query = query.Where("kategori_id = ?", kategori)
	}

	if status != "" {
		query = query.Where("LOWER(status) = ?", status)
	}

	if sort == "pending" || sort == "approved" || sort == "rejected" {
		query = query.Where("LOWER(status) = ?", sort)
	}

	if sort == "terbaru" {
		query = query.Order("tanggal_diajukan DESC")
	}

	if sort == "terlama" {
		query = query.Order("tanggal_diajukan ASC")
	}

	err := query.
		Limit(limit).
		Offset(offset).
		Find(&kegiatan).Error

	return kegiatan, err
}

func GetKegiatanSayaUsecase(userID uint, limit int, offset int, search string, kategori string, sort string, status string) ([]entities.Kegiatan, error) {
	var kegiatan []entities.Kegiatan

	query := database.DB.
		Preload("User").
		Preload("Kategori").
		Preload("FotoKegiatan").
		Where("user_id = ?", userID)

	if search != "" {
		query = query.Where("judul LIKE ?", "%"+search+"%")
	}

	if kategori != "" {
		query = query.Where("kategori_id = ?", kategori)
	}

	if status != "" {
		query = query.Where("LOWER(status) = ?", status)
	}

	if sort == "pending" || sort == "approved" || sort == "rejected" {
		query = query.Where("LOWER(status) = ?", sort)
	}

	if sort == "terbaru" {
		query = query.Order("tanggal_diajukan DESC")
	}

	if sort == "terlama" {
		query = query.Order("tanggal_diajukan ASC")
	}

	err := query.
		Limit(limit).
		Offset(offset).
		Find(&kegiatan).Error

	return kegiatan, err
}
