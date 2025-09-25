package repository

import (
	"evermos_rakamin/internal/entity"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type AlamatRepository interface {
	FindByUserID(ctx context.Context, userID int64) ([]entity.Alamat, error)
	FindAll(ctx context.Context) ([]entity.Alamat, error)
	FindByID(ctx context.Context, id int64) (*entity.Alamat, error)
	Update(ctx context.Context, user *entity.Alamat) error
	Create(ctx context.Context, user *entity.Alamat) error
	Delete(ctx context.Context, id int64) error
}

type alamatRepository struct {
	db *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) AlamatRepository {
	return &alamatRepository{db}
}



func (r *alamatRepository) FindAll(ctx context.Context) ([]entity.Alamat, error) {
	users := make([]entity.Alamat, 0)

	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *alamatRepository) FindByUserID(ctx context.Context, userID int64) ([]entity.Alamat, error) {
	var alamatList []entity.Alamat
	if err := r.db.WithContext(ctx).Where("id_user = ?", userID).Find(&alamatList).Error; err != nil {
		return nil, err
	}
	return alamatList, nil
}

func (r *alamatRepository) FindByID(ctx context.Context, id int64) (*entity.Alamat, error) {
	user := new(entity.Alamat)

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *alamatRepository) Create(ctx context.Context, user *entity.Alamat) error {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *alamatRepository) Update(ctx context.Context, user *entity.Alamat) error {
	if err := r.db.WithContext(ctx).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *alamatRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Model(&entity.Alamat{}).Delete("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}