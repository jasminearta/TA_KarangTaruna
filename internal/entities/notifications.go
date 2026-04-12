package entities

type Notification struct {
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Message string `json:"message"`
	IsRead  bool   `json:"is_read" gorm:"default:false"`

	User User `gorm:"foreignKey:UserID"`
}
