package entity

import "time"

type DetailTrx struct {
	ID         int64     `gorm:"primaryKey"`
	IdTrx      int64
	IdProduk   int64
	IdLogProduk   int64
	Kuantitas  int
	HargaTotal int
	UpdatedAt  time.Time
	CreatedAt  time.Time

	// Relasi
	Trx    Trx    `gorm:"foreignKey:IdTrx;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Produk Produk `gorm:"foreignKey:IdProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (DetailTrx) TableName() string {
	return "detail_trx"
}