package entities

type Kegiatan struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	UserID          uint   `json:"user_id"`
	KategoriID      uint   `json:"kategori_id"`
	Judul           string `json:"judul"`
	Deskripsi       string `json:"deskripsi"`
	TanggalBerjalan string `json:"tanggal_berjalan"`
	TanggalDiajukan string `json:"tanggal_diajukan"`
	Status          string `json:"status"`

	User         User           `gorm:"foreignKey:UserID" json:"user"`
	Kategori     Kategori       `gorm:"foreignKey:KategoriID" json:"kategori"`
	FotoKegiatan []FotoKegiatan `gorm:"foreignKey:KegiatanID" json:"foto_kegiatan"`
}

func (Kegiatan) TableName() string {
	return "kegiatan"
}
