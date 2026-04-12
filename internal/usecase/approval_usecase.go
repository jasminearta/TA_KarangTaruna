package usecase

import (
	"errors"
	"strings"
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func ApproveKegiatanUsecase(id uint, approvedBy uint) (entities.Kegiatan, error) {
	var kegiatan entities.Kegiatan

	tx := database.DB.Begin()

	err := tx.First(&kegiatan, id).Error
	if err != nil {
		tx.Rollback()
		return kegiatan, errors.New("kegiatan tidak ditemukan")
	}

	if strings.ToLower(kegiatan.Status) != "pending" {
		tx.Rollback()
		return kegiatan, errors.New("kegiatan sudah diproses")
	}

	kegiatan.Status = "approved"

	err = tx.Save(&kegiatan).Error
	if err != nil {
		tx.Rollback()
		return kegiatan, err
	}

	log := entities.ApprovalLog{
		ApprovedBy: approvedBy,
		Status:     "approved",
		Catatan:    "",
		TargetID:   kegiatan.ID,
		TargetType: "kegiatan",
	}

	err = tx.Create(&log).Error
	if err != nil {
		tx.Rollback()
		return kegiatan, err
	}

	tx.Commit()
	return kegiatan, nil
}

func RejectKegiatanUsecase(id uint, approvedBy uint, catatan string) (entities.Kegiatan, error) {
	var kegiatan entities.Kegiatan

	tx := database.DB.Begin()

	err := tx.First(&kegiatan, id).Error
	if err != nil {
		tx.Rollback()
		return kegiatan, errors.New("kegiatan tidak ditemukan")
	}

	if strings.ToLower(kegiatan.Status) != "pending" {
		tx.Rollback()
		return kegiatan, errors.New("kegiatan sudah diproses")
	}

	if strings.TrimSpace(catatan) == "" {
		tx.Rollback()
		return kegiatan, errors.New("catatan penolakan wajib diisi")
	}

	kegiatan.Status = "rejected"

	err = tx.Save(&kegiatan).Error
	if err != nil {
		tx.Rollback()
		return kegiatan, err
	}

	log := entities.ApprovalLog{
		ApprovedBy: approvedBy,
		Status:     "rejected",
		Catatan:    catatan,
		TargetID:   kegiatan.ID,
		TargetType: "kegiatan",
	}

	err = tx.Create(&log).Error
	if err != nil {
		tx.Rollback()
		return kegiatan, err
	}

	tx.Commit()
	return kegiatan, nil
}

func GetLatestApprovalLogByTarget(targetID uint, targetType string) (entities.ApprovalLog, error) {
	var log entities.ApprovalLog

	err := database.DB.
		Where("target_id = ? AND target_type = ?", targetID, targetType).
		Order("id DESC").
		First(&log).Error

	return log, err
}

func ApproveInovasiUsecase(id uint, approvedBy uint) (entities.Inovasi, error) {
	var inovasi entities.Inovasi

	tx := database.DB.Begin()

	err := tx.First(&inovasi, id).Error
	if err != nil {
		tx.Rollback()
		return inovasi, errors.New("inovasi tidak ditemukan")
	}

	if strings.ToLower(inovasi.Status) != "pending" {
		tx.Rollback()
		return inovasi, errors.New("inovasi sudah diproses")
	}

	inovasi.Status = "approved"

	err = tx.Save(&inovasi).Error
	if err != nil {
		tx.Rollback()
		return inovasi, err
	}

	log := entities.ApprovalLog{
		ApprovedBy: approvedBy,
		Status:     "approved",
		Catatan:    "",
		TargetID:   inovasi.ID,
		TargetType: "inovasi",
	}

	err = tx.Create(&log).Error
	if err != nil {
		tx.Rollback()
		return inovasi, err
	}

	tx.Commit()
	return inovasi, nil
}

func RejectInovasiUsecase(id uint, approvedBy uint, catatan string) (entities.Inovasi, error) {
	var inovasi entities.Inovasi

	tx := database.DB.Begin()

	err := tx.First(&inovasi, id).Error
	if err != nil {
		tx.Rollback()
		return inovasi, errors.New("inovasi tidak ditemukan")
	}

	if strings.ToLower(inovasi.Status) != "pending" {
		tx.Rollback()
		return inovasi, errors.New("inovasi sudah diproses")
	}

	if strings.TrimSpace(catatan) == "" {
		tx.Rollback()
		return inovasi, errors.New("catatan penolakan wajib diisi")
	}

	inovasi.Status = "rejected"

	err = tx.Save(&inovasi).Error
	if err != nil {
		tx.Rollback()
		return inovasi, err
	}

	log := entities.ApprovalLog{
		ApprovedBy: approvedBy,
		Status:     "rejected",
		Catatan:    catatan,
		TargetID:   inovasi.ID,
		TargetType: "inovasi",
	}

	err = tx.Create(&log).Error
	if err != nil {
		tx.Rollback()
		return inovasi, err
	}

	tx.Commit()
	return inovasi, nil
}
