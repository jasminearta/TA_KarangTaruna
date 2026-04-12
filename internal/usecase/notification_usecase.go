package usecase

import (
	"ta-karangtaruna/database"
	"ta-karangtaruna/internal/entities"
)

func CreateNotification(userID uint, title string, message string) error {

	notif := entities.Notification{
		UserID:  userID,
		Title:   title,
		Message: message,
		IsRead:  false,
	}

	return database.DB.Create(&notif).Error
}
