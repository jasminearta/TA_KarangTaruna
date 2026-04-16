package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

type DashboardData struct {
	TotalKegiatan    int64 `json:"total_kegiatan"`
	PendingKegiatan  int64 `json:"pending_kegiatan"`
	ApprovedKegiatan int64 `json:"approved_kegiatan"`
	RejectedKegiatan int64 `json:"rejected_kegiatan"`

	TotalInovasi    int64 `json:"total_inovasi"`
	PendingInovasi  int64 `json:"pending_inovasi"`
	ApprovedInovasi int64 `json:"approved_inovasi"`
	RejectedInovasi int64 `json:"rejected_inovasi"`

	TotalUsers int64 `json:"total_users"`
}

func GetDashboardData() (DashboardData, error) {
	var data DashboardData

	if err := database.DB.Model(&entities.Kegiatan{}).Count(&data.TotalKegiatan).Error; err != nil {
		return data, err
	}

	if err := database.DB.Model(&entities.Kegiatan{}).
		Where("LOWER(status) = ?", "pending").
		Count(&data.PendingKegiatan).Error; err != nil {
		return data, err
	}

	if err := database.DB.Model(&entities.Kegiatan{}).
		Where("LOWER(status) = ?", "approved").
		Count(&data.ApprovedKegiatan).Error; err != nil {
		return data, err
	}

	if err := database.DB.Model(&entities.Kegiatan{}).
		Where("LOWER(status) = ?", "rejected").
		Count(&data.RejectedKegiatan).Error; err != nil {
		return data, err
	}

	if err := database.DB.Model(&entities.Inovasi{}).Count(&data.TotalInovasi).Error; err != nil {
		return data, err
	}

	if err := database.DB.Model(&entities.Inovasi{}).
		Where("LOWER(status) = ?", "pending").
		Count(&data.PendingInovasi).Error; err != nil {
		return data, err
	}

	if err := database.DB.Model(&entities.Inovasi{}).
		Where("LOWER(status) = ?", "approved").
		Count(&data.ApprovedInovasi).Error; err != nil {
		return data, err
	}

	if err := database.DB.Model(&entities.Inovasi{}).
		Where("LOWER(status) = ?", "rejected").
		Count(&data.RejectedInovasi).Error; err != nil {
		return data, err
	}

	if err := database.DB.Model(&entities.User{}).Count(&data.TotalUsers).Error; err != nil {
		return data, err
	}

	return data, nil
}
