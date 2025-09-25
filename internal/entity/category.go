package entity

import "time"

type Category struct {
	ID           int64     `gorm:"primaryKey"`
	NamaCategory string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Relasi
	Produk []Produk `gorm:"foreignKey:IdCategory"`
	Logs   []LogProduk `gorm:"foreignKey:IdCategory"`
}

func (Category) TableName() string {
	return "category"
}