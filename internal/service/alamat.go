package service

import (
	"evermos_rakamin/internal/config"
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/entity"
	"evermos_rakamin/internal/repository"
	"context"
	"errors"
)

type AlamatService interface {
	GetAlamatById(ctx context.Context, id int64) (*dto.Alamat, error)
	GetAlamatByUserID(ctx context.Context, userID int64) ([]dto.Alamat, error)
	CreateAlamat(ctx context.Context, req *dto.CreateAlamatRequest) (*dto.Alamat, error)
	UpdateAlamat(ctx context.Context, req *dto.UpdateAlamatRequest) error
	DeleteAlamat(ctx context.Context, id int64) error
}

type alamatService struct {
	cfg        *config.Config
	alamatRepo repository.AlamatRepository
}

func NewAlamatService(cfg *config.Config, alamatRepo repository.AlamatRepository) AlamatService {
	return &alamatService{cfg, alamatRepo}
}

// Mapping helper
func mapAlamatToDTO(alamat *entity.Alamat) *dto.Alamat {
	return &dto.Alamat{
		ID:           alamat.ID,
		IdUser:       alamat.IdUser,
		JudulAlamat:  alamat.JudulAlamat,
		NamaPenerima: alamat.NamaPenerima,
		NoTelp:       alamat.NoTelp,
		DetailAlamat: alamat.DetailAlamat,
		CreatedAt:    alamat.CreatedAt,
		UpdatedAt:    alamat.UpdatedAt,
	}
}

func (s *alamatService) GetAlamatById(ctx context.Context, id int64) (*dto.Alamat, error) {
	alamat, err := s.alamatRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if alamat == nil {
		return nil, errors.New("alamat not found")
	}
	return mapAlamatToDTO(alamat), nil
}

func (s *alamatService) GetAlamatByUserID(ctx context.Context, userID int64) ([]dto.Alamat, error) {
	alamatList, err := s.alamatRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var result []dto.Alamat
	for _, alamat := range alamatList {
		result = append(result, *mapAlamatToDTO(&alamat))
	}
	return result, nil
}

func (s *alamatService) CreateAlamat(ctx context.Context, req *dto.CreateAlamatRequest) (*dto.Alamat, error) {
	alamat := &entity.Alamat{
		IdUser:       req.IdUser,
		JudulAlamat:  req.JudulAlamat,
		NamaPenerima: req.NamaPenerima,
		NoTelp:       req.NoTelp,
		DetailAlamat: req.DetailAlamat,
	}

	if err := s.alamatRepo.Create(ctx, alamat); err != nil {
		return nil, err
	}
	return mapAlamatToDTO(alamat), nil
}

func (s *alamatService) UpdateAlamat(ctx context.Context, req *dto.UpdateAlamatRequest) error {
	alamat, err := s.alamatRepo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if alamat == nil {
		return errors.New("alamat not found")
	}

	// Update hanya jika field-nya tidak nil
	if req.JudulAlamat != nil {
		alamat.JudulAlamat = *req.JudulAlamat
	}
	if req.NamaPenerima != nil {
		alamat.NamaPenerima = *req.NamaPenerima
	}
	if req.NoTelp != nil {
		alamat.NoTelp = *req.NoTelp
	}
	if req.DetailAlamat != nil {
		alamat.DetailAlamat = *req.DetailAlamat
	}

	return s.alamatRepo.Update(ctx, alamat)
}

func (s *alamatService) DeleteAlamat(ctx context.Context, id int64) error {
	return s.alamatRepo.Delete(ctx, id)
}