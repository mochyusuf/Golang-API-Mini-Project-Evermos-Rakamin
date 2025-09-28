package repository

import (
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/entity"
	"context"
    "fmt"

	"gorm.io/gorm"
)

type ProdukRepository interface {
	FindAll(ctx context.Context) ([]entity.Produk, error)
	FindByID(ctx context.Context, id int64) (*entity.Produk, error)
	FindByTokoID(ctx context.Context, tokoID int64) ([]entity.Produk, error)
	FindByCategoryID(ctx context.Context, categoryID int64) ([]entity.Produk, error)
	FindBySlug(ctx context.Context, slug string) (*entity.Produk, error)
	SearchByName(ctx context.Context, name string) ([]entity.Produk, error)
	Create(ctx context.Context, produk *entity.Produk) error
	Update(ctx context.Context, produk *entity.Produk) error
	Delete(ctx context.Context, id int64) error
	FindWithFilter(ctx context.Context, filter dto.ProdukFilterParams) ([]entity.Produk, error)
}

type produkRepository struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &produkRepository{db}
}

func (r *produkRepository) FindAll(ctx context.Context) ([]entity.Produk, error) {
	var produkList []entity.Produk
	if err := r.db.WithContext(ctx).Preload("Toko").Preload("Category").Preload("FotoProduk").Find(&produkList).Error; err != nil {
		return nil, err
	}
	return produkList, nil
}

func (r *produkRepository) FindByID(ctx context.Context, id int64) (*entity.Produk, error) {
	var produk entity.Produk
	if err := r.db.WithContext(ctx).Preload("Toko").Preload("Category").Preload("FotoProduk").
		Where("id = ?", id).First(&produk).Error; err != nil {
		return nil, err
	}
	return &produk, nil
}

func (r *produkRepository) FindByTokoID(ctx context.Context, tokoID int64) ([]entity.Produk, error) {
	var produkList []entity.Produk
	if err := r.db.WithContext(ctx).Where("id_toko = ?", tokoID).
		Preload("Category").Preload("FotoProduk").Find(&produkList).Error; err != nil {
		return nil, err
	}
	return produkList, nil
}

func (r *produkRepository) FindByCategoryID(ctx context.Context, categoryID int64) ([]entity.Produk, error) {
	var produkList []entity.Produk
	if err := r.db.WithContext(ctx).Where("id_category = ?", categoryID).
		Preload("Toko").Preload("FotoProduk").Find(&produkList).Error; err != nil {
		return nil, err
	}
	return produkList, nil
}

func (r *produkRepository) FindBySlug(ctx context.Context, slug string) (*entity.Produk, error) {
	var produk entity.Produk
	if err := r.db.WithContext(ctx).Where("slug = ?", slug).
		Preload("Toko").Preload("Category").Preload("FotoProduk").First(&produk).Error; err != nil {
		return nil, err
	}
	return &produk, nil
}

func (r *produkRepository) SearchByName(ctx context.Context, name string) ([]entity.Produk, error) {
	var produkList []entity.Produk
	if err := r.db.WithContext(ctx).Where("nama_produk LIKE ?", "%"+name+"%").
		Preload("Toko").Preload("Category").Preload("FotoProduk").Find(&produkList).Error; err != nil {
		return nil, err
	}
	return produkList, nil
}

func (r *produkRepository) Create(ctx context.Context, produk *entity.Produk) error {
	if err := r.db.WithContext(ctx).Create(produk).Error; err != nil {
		return err
	}
	return nil
}

func (r *produkRepository) Update(ctx context.Context, produk *entity.Produk) error {
	if err := r.db.WithContext(ctx).Where("id = ?", produk.ID).Updates(produk).Error; err != nil {
		return err
	}
	return nil
}

func (r *produkRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Produk{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *produkRepository) FindWithFilter(ctx context.Context, filter dto.ProdukFilterParams) ([]entity.Produk, error) {
	var produks []entity.Produk
	db := r.db.WithContext(ctx).Preload("Toko").Preload("Category").Preload("FotoProduk")

	if filter.NamaProduk != "" {
		db = db.Where("nama_produk LIKE ?", "%"+filter.NamaProduk+"%")
	}
	if filter.CategoryID != 0 {
		db = db.Where("id_category = ?", filter.CategoryID)
	}
	if filter.TokoID != 0 {
		db = db.Where("id_toko = ?", filter.TokoID)
	}
	if filter.MinHarga != 0 {
		db = db.Where("harga_konsumen >= ?", filter.MinHarga)
	}
	if filter.MaxHarga != 0 {
		db = db.Where("harga_konsumen <= ?", filter.MaxHarga)
	}

	offset := (filter.Page - 1) * filter.Limit
	if err := db.Limit(filter.Limit).Offset(offset).Find(&produks).Error; err != nil {
		return nil, err
	}
	return produks, nil
}