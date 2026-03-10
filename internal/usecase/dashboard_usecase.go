package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

type DashboardData struct {
	TotalKegiatan int64 `json:"total_kegiatan"`
	Pending       int64 `json:"pending"`
	Approved      int64 `json:"approved"`
	Rejected      int64 `json:"rejected"`
	TotalUsers    int64 `json:"total_users"`
}

func GetDashboardData() (DashboardData, error) {

	var data DashboardData

	database.DB.Model(&entities.Kegiatan{}).Count(&data.TotalKegiatan)

	database.DB.Model(&entities.Kegiatan{}).
		Where("status = ?", "pending").
		Count(&data.Pending)

	database.DB.Model(&entities.Kegiatan{}).
		Where("status = ?", "approved").
		Count(&data.Approved)

	database.DB.Model(&entities.Kegiatan{}).
		Where("status = ?", "rejected").
		Count(&data.Rejected)

	database.DB.Model(&entities.User{}).Count(&data.TotalUsers)

	return data, nil
}
