package entities

type FotoKegiatan struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	KegiatanID uint   `json:"kegiatan_id"`
	FileURL    string `json:"file_url"`

	Kegiatan Kegiatan `gorm:"foreignKey:KegiatanID" json:"-"`
}

func (FotoKegiatan) TableName() string {
	return "foto_kegiatan"
}
