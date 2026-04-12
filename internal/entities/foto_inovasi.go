package entities

type FotoInovasi struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	InovasiID uint   `json:"inovasi_id"`
	FileURL   string `json:"file_url"`

	Inovasi Inovasi `gorm:"foreignKey:InovasiID" json:"-"`
}

func (FotoInovasi) TableName() string {
	return "foto_inovasi"
}
