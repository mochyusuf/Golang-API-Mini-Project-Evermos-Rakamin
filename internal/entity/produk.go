package entity

import "time"

type Produk struct {
	ID            int64     `gorm:"primaryKey"`
	NamaProduk    string
	Slug          string
	HargaReseller int
	HargaKonsumen int
	Stok          int
	Deskripsi     string
	UpdatedAt     time.Time
	CreatedAt     time.Time
	IdToko        int64
	IdCategory    int64

	// Relasi
	Toko       Toko       `gorm:"foreignKey:IdToko;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category   Category   `gorm:"foreignKey:IdCategory;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FotoProduk []FotoProduk `gorm:"foreignKey:IdProduk"`
	Logs       []LogProduk  `gorm:"foreignKey:IdProduk"`
	DetailTrx  []DetailTrx  `gorm:"foreignKey:IdProduk"`
}

func (Produk) TableName() string {
	return "produk"
}