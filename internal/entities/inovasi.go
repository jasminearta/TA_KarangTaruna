package entities

type Inovasi struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	UserID          uint   `json:"user_id"`
	KategoriID      uint   `json:"kategori_id"`
	Judul           string `json:"judul"`
	Deskripsi       string `json:"deskripsi"`
	TanggalDiajukan string `json:"tanggal_diajukan"`
	Status          string `json:"status"`

	User        User          `gorm:"foreignKey:UserID" json:"user"`
	Kategori    Kategori      `gorm:"foreignKey:KategoriID" json:"kategori"`
	FotoInovasi []FotoInovasi `gorm:"foreignKey:InovasiID" json:"foto_inovasi"`
}

func (Inovasi) TableName() string {
	return "inovasi"
}
