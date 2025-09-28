package service

import (
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/entity"
	"evermos_rakamin/internal/repository"
	"context"
	"errors"
)

type ProdukService interface {
	GetAllProduk(ctx context.Context, filter dto.ProdukFilterParams) (*dto.ProdukListPaginated, error)
	GetProdukByID(ctx context.Context, id int64) (*dto.ProdukResponse, error)
	CreateProduk(ctx context.Context, req *dto.CreateProdukRequest,  urlFotos []string) error
	UpdateProduk(ctx context.Context, id int64, req *dto.UpdateProdukRequest) error
	DeleteProduk(ctx context.Context, id int64) error
}

type produkService struct {
	produkRepo 			repository.ProdukRepository
	fotoProdukRepo   	repository.FotoProdukRepository
}

func NewProdukService(produkRepo repository.ProdukRepository, fotoProdukRepo repository.FotoProdukRepository) ProdukService {
	return &produkService{produkRepo, fotoProdukRepo}
}

func (s *produkService) GetAllProduk(ctx context.Context, filter dto.ProdukFilterParams) (*dto.ProdukListPaginated, error) {
	produks, err := s.produkRepo.FindWithFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ProdukResponse, len(produks))
	for i, p := range produks {
		photos := make([]dto.PhotoResponse, len(p.FotoProduk))
		for j, photo := range p.FotoProduk {
			photos[j] = dto.PhotoResponse{
				ID:        photo.ID,
				ProductID: photo.IdProduk,
				Url:       photo.Url,
			}
		}

		responses[i] = dto.ProdukResponse{
			ID:            p.ID,
			NamaProduk:    p.NamaProduk,
			Slug:          p.Slug,
			HargaReseller: p.HargaReseller,
			HargaKonsumen: p.HargaKonsumen,
			Stok:          p.Stok,
			Deskripsi:     p.Deskripsi,
			Toko: dto.TokoResponse{
				ID:       p.Toko.ID,
				NamaToko: p.Toko.NamaToko,
				UrlFoto:  p.Toko.UrlFoto,
			},
			Category: dto.CategoryResponse{
				ID:           p.Category.ID,
				NamaCategory: p.Category.NamaCategory,
			},
			Photos: photos,
		}
	}

	return &dto.ProdukListPaginated{
		Data:  responses,
		Page:  filter.Page,
		Limit: filter.Limit,
	}, nil
}

func (s *produkService) GetProdukByID(ctx context.Context, id int64) (*dto.ProdukResponse, error) {
	produk, err := s.produkRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if produk == nil {
		return nil, errors.New("produk not found")
	}
	return &dto.ProdukResponse{
		ID:            produk.ID,
		NamaProduk:    produk.NamaProduk,
		Slug:          produk.Slug,
		HargaReseller: produk.HargaReseller,
		HargaKonsumen: produk.HargaKonsumen,
		Stok:          produk.Stok,
		Deskripsi:     produk.Deskripsi,
		Toko: dto.TokoResponse{
			ID:       produk.Toko.ID,
			NamaToko: produk.Toko.NamaToko,
			UrlFoto:  produk.Toko.UrlFoto,
		},
		Category: dto.CategoryResponse{
			ID:           produk.Category.ID,
			NamaCategory: produk.Category.NamaCategory,
		},

	}, nil
}

func (s *produkService) CreateProduk(ctx context.Context, req *dto.CreateProdukRequest,  urlFotos []string) error {
	produk := &entity.Produk{
		NamaProduk:    req.NamaProduk,
		Slug:          req.Slug,
		HargaReseller: req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
		IdToko:        req.IdToko,
		IdCategory:    req.IdCategory,
	}
	if err := s.produkRepo.Create(ctx, produk); err != nil {
		return err
	}

	var fotoProduks []entity.FotoProduk
	for i := 0; i < len(urlFotos); i++ {
		fotoProduks = append(fotoProduks, entity.FotoProduk{
			IdProduk: produk.ID,
			Url:      urlFotos[i],
		})
	}

	if err := s.fotoProdukRepo.CreateBulk(ctx, fotoProduks); err != nil {
		return err
	}


	return nil
}

func (s *produkService) UpdateProduk(ctx context.Context, id int64, req *dto.UpdateProdukRequest) error {
	produk, err := s.produkRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if produk == nil {
		return errors.New("produk not found")
	}

	if req.NamaProduk != nil {
		produk.NamaProduk = *req.NamaProduk
	}
	if req.Slug != nil {
		produk.Slug = *req.Slug
	}
	if req.HargaReseller != nil {
		produk.HargaReseller = *req.HargaReseller
	}
	if req.HargaKonsumen != nil {
		produk.HargaKonsumen = *req.HargaKonsumen
	}
	if req.Stok != nil {
		produk.Stok = *req.Stok
	}
	if req.Deskripsi != nil {
		produk.Deskripsi = *req.Deskripsi
	}
	if req.IdCategory != nil {
		produk.IdCategory = *req.IdCategory
	}

	return s.produkRepo.Update(ctx, produk)
}

func (s *produkService) DeleteProduk(ctx context.Context, id int64) error {
	return s.produkRepo.Delete(ctx, id)
}