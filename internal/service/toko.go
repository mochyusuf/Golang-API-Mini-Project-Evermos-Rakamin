package service

import (
	"context"
	"errors"

	"evermos_rakamin/internal/config"
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/entity"
	"evermos_rakamin/internal/repository"
)

type TokoService interface {
	GetAllToko(ctx context.Context) ([]dto.TokoResponse, error)
	GetTokoByID(ctx context.Context, id int64) (*dto.TokoResponse, error)
	GetTokoByUserID(ctx context.Context, userID int64) (*dto.TokoResponse, error)
	SearchTokoByName(ctx context.Context, name string) ([]dto.TokoResponse, error)
	CreateToko(ctx context.Context, req *dto.CreateTokoRequest) (*dto.TokoResponse, error)
	UpdateToko(ctx context.Context, id int64,req *dto.UpdateTokoRequest) error
	DeleteToko(ctx context.Context, id int64) error
	GetTokoPaginated(ctx context.Context, limit, page int, namaToko string) (*PaginatedTokoResponse, error)
}

type tokoService struct {
	cfg         *config.Config
	tokoRepo    repository.TokoRepository
}

func NewTokoService(cfg *config.Config, tokoRepo repository.TokoRepository) TokoService {
	return &tokoService{cfg, tokoRepo}
}

func mapTokoToResponse(toko *entity.Toko) *dto.TokoResponse {
	return &dto.TokoResponse{
		ID:        toko.ID,
		IdUser:    toko.IdUser,
		NamaToko:  toko.NamaToko,
		UrlFoto:   toko.UrlFoto,
		CreatedAt: toko.CreatedAt,
		UpdatedAt: toko.UpdatedAt,
	}
}

type PaginatedTokoResponse struct {
	Data       []dto.TokoResponse `json:"data"`
	TotalItems int64              `json:"total_items"`
	Page       int                `json:"page"`
	Limit      int                `json:"limit"`
	TotalPages int                `json:"total_pages"`
}

func (s *tokoService) GetAllToko(ctx context.Context) ([]dto.TokoResponse, error) {
	tokos, err := s.tokoRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.TokoResponse
	for _, toko := range tokos {
		responses = append(responses, *mapTokoToResponse(&toko))
	}
	return responses, nil
}

func (s *tokoService) GetTokoByID(ctx context.Context, id int64) (*dto.TokoResponse, error) {
	toko, err := s.tokoRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if toko == nil {
		return nil, errors.New("toko not found")
	}
	return mapTokoToResponse(toko), nil
}

func (s *tokoService) GetTokoByUserID(ctx context.Context, userID int64) (*dto.TokoResponse, error) {
	toko, err := s.tokoRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if toko == nil {
		return nil, errors.New("toko not found")
	}
	return mapTokoToResponse(toko), nil
}

func (s *tokoService) SearchTokoByName(ctx context.Context, name string) ([]dto.TokoResponse, error) {
	tokos, err := s.tokoRepo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	var responses []dto.TokoResponse
	for _, toko := range tokos {
		responses = append(responses, *mapTokoToResponse(&toko))
	}
	return responses, nil
}

func (s *tokoService) CreateToko(ctx context.Context, req *dto.CreateTokoRequest) (*dto.TokoResponse, error) {
	toko := &entity.Toko{
		IdUser:    req.IdUser,

	}

	err := s.tokoRepo.Create(ctx, toko)
	if err != nil {
		return nil, err
	}
	return mapTokoToResponse(toko), nil
}

func (s *tokoService) UpdateToko(ctx context.Context, id int64, req *dto.UpdateTokoRequest) error {
	toko, err := s.tokoRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if toko == nil {
		return errors.New("toko not found")
	}

	if req.NamaToko != nil {
		toko.NamaToko = *req.NamaToko
	}
	if req.UrlFoto != nil {
		toko.UrlFoto = *req.UrlFoto
	}

	return s.tokoRepo.Update(ctx, toko)
}

func (s *tokoService) DeleteToko(ctx context.Context, id int64) error {
	return s.tokoRepo.Delete(ctx, id)
}

func (s *tokoService) GetTokoPaginated(ctx context.Context, limit, page int, namaToko string) (*PaginatedTokoResponse, error) {
	offset := (page - 1) * limit

	tokos, err := s.tokoRepo.FindPaginated(ctx, limit, offset, namaToko)
	if err != nil {
		return nil, err
	}

	totalItems, err := s.tokoRepo.CountFiltered(ctx, namaToko)
	if err != nil {
		return nil, err
	}

	var responses []dto.TokoResponse
	for _, toko := range tokos {
		responses = append(responses, *mapTokoToResponse(&toko))
	}

	totalPages := int((totalItems + int64(limit) - 1) / int64(limit)) // ceil division

	return &PaginatedTokoResponse{
		Data:       responses,
		TotalItems: totalItems,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}