package repository

import (
	"evermos_rakamin/internal/entity"
	"context"

	"gorm.io/gorm"
)

type TokoRepository interface {
    FindByUserID(ctx context.Context, userID int64) (*entity.Toko, error)
    FindAll(ctx context.Context) ([]entity.Toko, error)
    FindByID(ctx context.Context, id int64) (*entity.Toko, error)
    FindByName(ctx context.Context, namaToko string) ([]entity.Toko, error)
    Update(ctx context.Context, toko *entity.Toko) error
    Create(ctx context.Context, toko *entity.Toko) error
    Delete(ctx context.Context, id int64) error
	FindPaginated(ctx context.Context, limit, offset int, namaToko string) ([]entity.Toko, error)
	CountFiltered(ctx context.Context, namaToko string) (int64, error)
}

type tokoRepository struct {
    db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
    return &tokoRepository{db}
}

func (r *tokoRepository) FindAll(ctx context.Context) ([]entity.Toko, error) {
    tokos := make([]entity.Toko, 0)
    if err := r.db.WithContext(ctx).Find(&tokos).Error; err != nil {
        return nil, err
    }
    return tokos, nil
}

func (r *tokoRepository) FindByUserID(ctx context.Context, userID int64) (*entity.Toko, error) {
    toko := new(entity.Toko)
    if err := r.db.WithContext(ctx).Where("id_user = ?", userID).First(&toko).Error; err != nil {
        return nil, err
    }
    return toko, nil
}

func (r *tokoRepository) FindByID(ctx context.Context, id int64) (*entity.Toko, error) {
    toko := new(entity.Toko)
    if err := r.db.WithContext(ctx).Where("id = ?", id).First(&toko).Error; err != nil {
        return nil, err
    }
    return toko, nil
}

func (r *tokoRepository) FindByName(ctx context.Context, namaToko string) ([]entity.Toko, error) {
    var tokos []entity.Toko
    if err := r.db.WithContext(ctx).Where("nama_toko LIKE ?", "%"+namaToko+"%").Find(&tokos).Error; err != nil {
        return nil, err
    }
    return tokos, nil
}


func (r *tokoRepository) Create(ctx context.Context, toko *entity.Toko) error {
    if err := r.db.WithContext(ctx).Create(&toko).Error; err != nil {
        return err
    }
    return nil
}

func (r *tokoRepository) Update(ctx context.Context, toko *entity.Toko) error {
    if err := r.db.WithContext(ctx).Where("id = ?", toko.ID).Updates(&toko).Error; err != nil {
        return err
    }
    return nil
}

func (r *tokoRepository) Delete(ctx context.Context, id int64) error {
    if err := r.db.WithContext(ctx).Delete(&entity.Toko{}, "id = ?", id).Error; err != nil {
        return err
    }
    return nil
}

func (r *tokoRepository) FindPaginated(ctx context.Context, limit, offset int, namaToko string) ([]entity.Toko, error) {
	var tokos []entity.Toko
	query := r.db.WithContext(ctx)

	if namaToko != "" {
		query = query.Where("nama_toko LIKE ?", "%"+namaToko+"%")
	}

	if err := query.Limit(limit).Offset(offset).Find(&tokos).Error; err != nil {
		return nil, err
	}
	return tokos, nil
}

func (r *tokoRepository) CountFiltered(ctx context.Context, namaToko string) (int64, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&entity.Toko{})

	if namaToko != "" {
		query = query.Where("nama_toko LIKE ?", "%"+namaToko+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}