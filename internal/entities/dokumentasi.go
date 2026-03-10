package entities

type Dokumentasi struct {
	ID         uint   `gorm:"primaryKey"`
	KegiatanID uint   `json:"kegiatan_id"`
	FileURL    string `json:"file_url"`

	Kegiatan Kegiatan `gorm:"foreignKey:KegiatanID"`
}

func (Dokumentasi) TableName() string {
	return "dokumentasi"
}
