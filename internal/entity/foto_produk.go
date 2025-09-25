package entity

import "time"

type FotoProduk struct {
	ID        int64     `gorm:"primaryKey"`
	IdProduk  int64
	Url       string
	UpdatedAt time.Time
	CreatedAt time.Time

	// Relasi
	Produk Produk `gorm:"foreignKey:IdProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (FotoProduk) TableName() string {
	return "foto_produk"
}