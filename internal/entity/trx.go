package entity

import "time"

type Trx struct {
	ID               int64     `gorm:"primaryKey"`
	IdUser           int64
	AlamatPengiriman int64
	HargaTotal       int
	KodeInvoice      string
	MethodBayar      string
	UpdatedAt        time.Time
	CreatedAt        time.Time

	// Relasi
	User   User    `gorm:"foreignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Alamat Alamat  `gorm:"foreignKey:AlamatPengiriman;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Detail []DetailTrx `gorm:"foreignKey:IdTrx"`
}

func (Trx) TableName() string {
	return "trx"
}