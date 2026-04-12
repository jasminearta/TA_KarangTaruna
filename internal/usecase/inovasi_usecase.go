package usecase

import (
	"errors"
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func CreateInovasi(
	judul string,
	deskripsi string,
	tanggalDiajukan string,
	kategoriID uint,
	userID uint,
) (entities.Inovasi, error) {

	inovasi := entities.Inovasi{
		Judul:           judul,
		Deskripsi:       deskripsi,
		TanggalDiajukan: tanggalDiajukan,
		KategoriID:      kategoriID,
		UserID:          userID,
		Status:          "pending",
	}

	err := database.DB.Create(&inovasi).Error
	if err != nil {
		return inovasi, err
	}

	database.DB.
		Preload("User").
		Preload("Kategori").
		First(&inovasi, inovasi.ID)

	return inovasi, nil
}

func GetInovasi() ([]entities.Inovasi, error) {
	var inovasi []entities.Inovasi

	err := database.DB.
		Preload("User").
		Preload("Kategori").
		Where("status = ?", "approved").
		Find(&inovasi).Error

	return inovasi, err
}

func GetAllInovasi(limit int, offset int, search string, kategori string, sort string, status string) ([]entities.Inovasi, error) {
	var inovasi []entities.Inovasi

	query := database.DB.
		Preload("User").
		Preload("Kategori").
		Preload("FotoInovasi")

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
		Find(&inovasi).Error

	return inovasi, err
}

func GetInovasiByUser(userID uint) ([]entities.Inovasi, error) {
	var inovasi []entities.Inovasi

	err := database.DB.
		Preload("User").
		Preload("Kategori").
		Preload("FotoInovasi").
		Where("user_id = ?", userID).
		Find(&inovasi).Error
	return inovasi, err
}

func GetDetailInovasi(id uint) (entities.Inovasi, error) {
	var inovasi entities.Inovasi

	err := database.DB.
		Preload("User").
		Preload("Kategori").
		Where("status = ?", "approved").
		First(&inovasi, id).Error

	return inovasi, err
}

func UpdateInovasi(
	id uint,
	judul string,
	deskripsi string,
	tanggalDiajukan string,
	kategoriID uint,
	userID uint,
) (entities.Inovasi, error) {

	var inovasi entities.Inovasi

	err := database.DB.First(&inovasi, id).Error
	if err != nil {
		return inovasi, err
	}

	if inovasi.UserID != userID {
		return inovasi, errors.New("anda tidak memiliki akses untuk mengedit inovasi ini")
	}

	// sesuai aturan sekarang: hanya yang masih pending yang boleh diubah data utamanya
	if inovasi.Status != "pending" {
		return inovasi, errors.New("inovasi yang sudah diverifikasi tidak dapat diubah")
	}

	inovasi.Judul = judul
	inovasi.Deskripsi = deskripsi
	inovasi.TanggalDiajukan = tanggalDiajukan
	inovasi.KategoriID = kategoriID

	err = database.DB.Save(&inovasi).Error
	if err != nil {
		return inovasi, err
	}

	err = database.DB.
		Preload("User").
		Preload("Kategori").
		Preload("FotoInovasi").
		First(&inovasi, inovasi.ID).Error

	return inovasi, err
}

func DeleteInovasi(id uint, userID uint) error {
	var inovasi entities.Inovasi

	err := database.DB.First(&inovasi, id).Error
	if err != nil {
		return errors.New("inovasi tidak ditemukan")
	}

	if inovasi.UserID != userID {
		return errors.New("anda tidak memiliki akses")
	}

	if inovasi.Status != "pending" {
		return errors.New("inovasi yang sudah diverifikasi tidak bisa dihapus")
	}

	if err := database.DB.Delete(&inovasi).Error; err != nil {
		return err
	}

	return nil
}

func GetAllInovasiKetuaUsecase(limit int, offset int, search string, kategori string, sort string, status string) ([]entities.Inovasi, error) {
	var inovasi []entities.Inovasi

	query := database.DB.
		Preload("User").
		Preload("Kategori").
		Preload("FotoInovasi")

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
		Find(&inovasi).Error

	return inovasi, err
}

func GetInovasiSayaUsecase(userID uint, limit int, offset int, search string, kategori string, sort string, status string) ([]entities.Inovasi, error) {
	var inovasi []entities.Inovasi

	query := database.DB.
		Preload("User").
		Preload("Kategori").
		Preload("FotoInovasi").
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
		Find(&inovasi).Error

	return inovasi, err
}
