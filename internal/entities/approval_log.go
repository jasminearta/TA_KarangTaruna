package entities

type ApprovalLog struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ApprovedBy uint   `gorm:"not null" json:"approved_by"`
	Status     string `gorm:"type:enum('approved','rejected')" json:"status"`
	Catatan    string `gorm:"type:text" json:"catatan"`
	TargetID   uint   `gorm:"not null" json:"target_id"`
	TargetType string `gorm:"type:varchar(50);not null" json:"target_type"`

	User User `gorm:"foreignKey:ApprovedBy" json:"user"`
}

func (ApprovalLog) TableName() string {
	return "approval_log"
}
