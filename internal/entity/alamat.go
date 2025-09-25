package entity

import "time"

type Alamat struct {
	ID           int64     `gorm:"primaryKey"`
	IdUser       int64
	JudulAlamat  string
	NamaPenerima string
	NoTelp       string
	DetailAlamat string
	UpdatedAt    time.Time
	CreatedAt    time.Time

	// Relasi
	User User `gorm:"foreignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Alamat) TableName() string {
	return "alamat"
}