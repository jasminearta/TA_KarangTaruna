package entities

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Foto     string `json:"foto"`
	Role     string `json:"role"`
}

func (User) TableName() string {
	return "user"
}
