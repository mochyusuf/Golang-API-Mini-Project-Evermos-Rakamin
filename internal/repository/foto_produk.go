package repository

import (
	"evermos_rakamin/internal/entity"
	"context"

	"gorm.io/gorm"
)

type fotoProdukRepository struct {
	db *gorm.DB
}

type FotoProdukRepository interface {
	CreateBulk(ctx context.Context, fotos []entity.FotoProduk) error
	FindByProdukID(ctx context.Context, produkID int64) ([]entity.FotoProduk, error)
	DeleteByID(ctx context.Context, id int64) error
}



func NewFotoProdukRepository(db *gorm.DB) FotoProdukRepository {
	return &fotoProdukRepository{db: db}
}

func (r *fotoProdukRepository) CreateBulk(ctx context.Context, fotos []entity.FotoProduk) error {
	return r.db.WithContext(ctx).Create(&fotos).Error
}

func (r *fotoProdukRepository) FindByProdukID(ctx context.Context, produkID int64) ([]entity.FotoProduk, error) {
	var fotos []entity.FotoProduk
	err := r.db.WithContext(ctx).Where("id_produk = ?", produkID).Find(&fotos).Error
	if err != nil {
		return nil, err
	}
	return fotos, nil
}

func (r *fotoProdukRepository) DeleteByID(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.FotoProduk{}, id).Error
}