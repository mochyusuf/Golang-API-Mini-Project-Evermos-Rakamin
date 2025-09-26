package repository

import (
	"evermos_rakamin/internal/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]entity.Category, error)
	FindByID(ctx context.Context, id int64) (*entity.Category, error)
	FindByName(ctx context.Context, name string) ([]entity.Category, error)
	Create(ctx context.Context, category *entity.Category) error
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id int64) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) FindAll(ctx context.Context) ([]entity.Category, error) {
	var categories []entity.Category
	if err := r.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) FindByID(ctx context.Context, id int64) (*entity.Category, error) {
	var category entity.Category
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) FindByName(ctx context.Context, name string) ([]entity.Category, error) {
	var categories []entity.Category
	if err := r.db.WithContext(ctx).Where("nama_category LIKE ?", "%"+name+"%").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) Create(ctx context.Context, category *entity.Category) error {
	if err := r.db.WithContext(ctx).Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) Update(ctx context.Context, category *entity.Category) error {
	if err := r.db.WithContext(ctx).Where("id = ?", category.ID).Updates(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}