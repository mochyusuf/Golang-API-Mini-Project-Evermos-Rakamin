package entity

import "time"

type Toko struct {
	ID         int64     `gorm:"primaryKey"`
	IdUser     int64
	NamaToko   string
	UrlFoto    string
	UpdatedAt  time.Time
	CreatedAt  time.Time

	// Relasi
	User    User     `gorm:"foreignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Produk  []Produk `gorm:"foreignKey:IdToko"`
	Logs    []LogProduk `gorm:"foreignKey:IdToko"`
}

func (Toko) TableName() string {
	return "toko"
}