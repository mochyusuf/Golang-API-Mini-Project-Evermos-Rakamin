package entity

import "time"

type LogProduk struct {
	ID            int64     `gorm:"primaryKey"`
	IdProduk      int64
	NamaProduk    string
	Slug          string
	HargaReseller int
	HargaKonsumen int
	Deskripsi     string
	CreatedAt     time.Time
	IdToko        int64
	IdCategory    int64

	// Relasi
	Produk   Produk   `gorm:"foreignKey:IdProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Toko     Toko     `gorm:"foreignKey:IdToko;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category Category `gorm:"foreignKey:IdCategory;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (LogProduk) TableName() string {
	return "log_produk"
}