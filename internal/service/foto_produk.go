package service

import (
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/entity"
	"evermos_rakamin/internal/repository"
	"context"
)

type FotoProdukService interface {
	CreateFotoProduk(ctx context.Context, req *dto.FotoProdukReq) error
}

type fotoProdukService struct {
	fotoProdukRepository repository.FotoProdukRepository
}

func NewFotoProdukService(fotoProdukRepository repository.FotoProdukRepository) FotoProdukService {
	return &fotoProdukService{fotoProdukRepository}
}


func (s *fotoProdukService) CreateFotoProduk(ctx context.Context, req *dto.FotoProdukReq) error {
	fotoProduk := &entity.FotoProduk{
		IdProduk: req.IdProduk,
		Url:      req.Url,
	}

	if err := s.fotoProdukRepository.CreateBulk(ctx, []entity.FotoProduk{*fotoProduk}); err != nil {
		return err
	}

	return nil
}