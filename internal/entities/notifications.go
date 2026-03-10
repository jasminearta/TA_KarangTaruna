package entities

import "time"

type Notification struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Title     string `gorm:"type:varchar(255)"`
	Message   string `gorm:"type:text"`
	IsRead    bool   `gorm:"default:false"`
	CreatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}
