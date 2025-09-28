package repository

import (
	"evermos_rakamin/internal/entity"
	"context"

	"gorm.io/gorm"
)

type trxRepository struct {
	db *gorm.DB
}
type TrxRepository interface {
	GetAll(ctx context.Context, limit, offset int) ([]entity.Trx, error)
	FindByID(ctx context.Context, id int64) (*entity.Trx, error)
	Create(ctx context.Context, trx *entity.Trx) error
	GetProdukByID(ctx context.Context, id int64) (*entity.Produk, error)
}

func NewTrxRepository(db *gorm.DB) TrxRepository {
	return &trxRepository{db: db}
}

func (r *trxRepository) GetAll(ctx context.Context, limit, offset int) ([]entity.Trx, error) {
	var trxs []entity.Trx
	err := r.db.WithContext(ctx).
		Preload("Alamat").
		Preload("Detail.Produk.Toko").
		Preload("Detail.Produk.Category").
		Preload("Detail.Produk.FotoProduk").
		Limit(limit).
		Offset(offset).
		Find(&trxs).Error

	if err != nil {
		return nil, err
	}

	return trxs, nil
}

func (r *trxRepository) FindByID(ctx context.Context, id int64) (*entity.Trx, error) {
	var trx entity.Trx

	err := r.db.WithContext(ctx).
		Preload("Alamat").
		Preload("Detail.Produk.Toko").
		Preload("Detail.Produk.Category").
		Preload("Detail.Produk.FotoProduk").
		Where("id = ?", id).
		First(&trx).Error

	if err != nil {
		return nil, err
	}

	return &trx, nil
}
func (r *trxRepository) Create(ctx context.Context, trx *entity.Trx) error {
	return r.db.WithContext(ctx).Create(trx).Error
}

func (r *trxRepository) GetProdukByID(ctx context.Context, id int64) (*entity.Produk, error) {
	var produk entity.Produk
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&produk).Error; err != nil {
		return nil, err
	}
	return &produk, nil
}