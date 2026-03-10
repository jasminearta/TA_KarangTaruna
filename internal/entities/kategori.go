package entities

type Kategori struct {
	ID   uint   `gorm:"primaryKey"`
	Nama string `gorm:"type:varchar(100);not null"`
}

func (Kategori) TableName() string {
	return "kategori"
}
