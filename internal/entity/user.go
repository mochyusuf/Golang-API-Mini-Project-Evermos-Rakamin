package entity

import "time"

type User struct {
	ID            int64     `gorm:"primaryKey"`
	Nama          string
	KataSandi     string
	NoTelp        string    `gorm:"unique"`
	TanggalLahir  time.Time
	JenisKelamin  string
	Tentang       string
	Pekerjaan     string
	Email         string	`gorm:"unique"`
	IdProvinsi    string
	IdKota        string
	IsAdmin	      bool
	UpdatedAt     time.Time
	CreatedAt     time.Time

	// Relasi
	Alamat []Alamat `gorm:"foreignKey:IdUser"`
	Toko   []Toko   `gorm:"foreignKey:IdUser"`
	Trx    []Trx    `gorm:"foreignKey:IdUser"`
}

func (User) TableName() string {
	return "user"
}