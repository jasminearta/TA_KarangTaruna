package entities

type ApprovalLog struct {
	ID         uint   `gorm:"primaryKey"`
	KegiatanID uint   `gorm:"not null"`
	ApprovedBy uint   `gorm:"not null"`
	Status     string `gorm:"type:enum('approved','rejected')"`
	Catatan    string `gorm:"type:text"`

	Kegiatan Kegiatan `gorm:"foreignKey:KegiatanID"`
	User     User     `gorm:"foreignKey:ApprovedBy"`
}

func (ApprovalLog) TableName() string {
	return "approval_log"
}
