package entities

import "time"

type Komentar struct {
	ID         uint   `gorm:"primaryKey"`
	KegiatanID uint   `gorm:"not null"`
	UserID     uint   `gorm:"not null"`
	Isi        string `gorm:"type:text"`
	CreatedAt  time.Time

	Kegiatan Kegiatan `gorm:"foreignKey:KegiatanID"`
	User     User     `gorm:"foreignKey:UserID"`
}

func (Komentar) TableName() string {
	return "komentar"
}
