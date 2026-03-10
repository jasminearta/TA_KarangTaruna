package entities

import (
	"time"
)

type Kegiatan struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"not null"`
	KategoriID uint   `gorm:"not null"`
	Judul      string `gorm:"type:varchar(255);not null"`
	Deskripsi  string `gorm:"type:text"`
	Tanggal    time.Time
	Status     string `gorm:"type:enum('Pending','Approved','Rejected');default:'Pending'"`

	User     User     `gorm:"foreignKey:UserID"`
	Kategori Kategori `gorm:"foreignKey:KategoriID"`
}

func (Kegiatan) TableName() string {
	return "kegiatan"
}
