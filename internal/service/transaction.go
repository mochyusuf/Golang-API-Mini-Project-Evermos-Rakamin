package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"evermos_rakamin/internal/config"
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/entity"
	"evermos_rakamin/internal/repository"
)

type TrxService interface {
	CreateTrx(ctx context.Context, req *dto.CreateTrxRequest, userID int64) error
	GetAllTrx(ctx context.Context, page, limit int) (*dto.TrxListPaginated, error)
	GetTrxByID(ctx context.Context, id int64) (*dto.TrxResponse, error)

}

type trxService struct {
	cfg      *config.Config
	trxRepo  repository.TrxRepository
}

func NewTrxService(cfg *config.Config, trxRepo repository.TrxRepository) TrxService {
	return &trxService{cfg, trxRepo}
}

func (s *trxService) CreateTrx(ctx context.Context, req *dto.CreateTrxRequest, userID int64) error {
	var totalHarga int

	// Generate Kode Invoice
	kodeInvoice := fmt.Sprintf("INV-%d", time.Now().Unix())

	// Ambil data produk + hitung harga total
	var detailTrxEntities []entity.DetailTrx
	for _, detail := range req.DetailTrx {
		// Ambil Produk (pastikan produk ada)
		produk, err := s.trxRepo.GetProdukByID(ctx, detail.ProductID)
		if err != nil {
			return fmt.Errorf("produk id %d not found", detail.ProductID)
		}

		hargaTotalProduk := produk.HargaKonsumen * detail.Kuantitas
		totalHarga += hargaTotalProduk

		detailTrxEntities = append(detailTrxEntities, entity.DetailTrx{
			IdProduk:   detail.ProductID,
			Kuantitas:  detail.Kuantitas,
			HargaTotal: hargaTotalProduk,
		})
	}

	// Buat entity trx
	trxEntity := &entity.Trx{
		IdUser:           userID,
		AlamatPengiriman: req.AlamatKirim,
		HargaTotal:       totalHarga,
		KodeInvoice:      kodeInvoice,
		MethodBayar:      req.MethodBayar,
		Detail:           detailTrxEntities, // akan auto insert detail_trx jika pakai GORM trx
	}

	// Simpan trx + detail_trx dalam 1 transaksi DB
	if err := s.trxRepo.Create(ctx, trxEntity); err != nil {
		return err
	}

	return nil
}


func (s *trxService) GetAllTrx(ctx context.Context, page, limit int) (*dto.TrxListPaginated, error) {
	offset := (page - 1) * limit

	// Get Data dari Repo
	trxs, err := s.trxRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	// Mapping Entity -> DTO Response
	var responses []dto.TrxResponse
	for _, trx := range trxs {
		response := dto.TrxResponse{
			ID:          trx.ID,
			HargaTotal:  trx.HargaTotal,
			KodeInvoice: trx.KodeInvoice,
			MethodBayar: trx.MethodBayar,
			AlamatKirim: dto.AlamatResponse{
				ID:            trx.Alamat.ID,
				JudulAlamat:   trx.Alamat.JudulAlamat,
				NamaPenerima:  trx.Alamat.NamaPenerima,
				NoTelp:        trx.Alamat.NoTelp,
				DetailAlamat:  trx.Alamat.DetailAlamat,
			},
		}

		for _, detail := range trx.Detail {
			// Photos Mapping
			photos := []dto.PhotoResponse{}
			for _, photo := range detail.Produk.FotoProduk {
				photos = append(photos, dto.PhotoResponse{
					ID:        photo.ID,
					ProductID: photo.IdProduk,
					Url:       photo.Url,
				})
			}

			detailResponse := dto.DetailTrxResponse{
				Product: dto.ProdukResponse{
					ID:            detail.Produk.ID,
					NamaProduk:    detail.Produk.NamaProduk,
					Slug:          detail.Produk.Slug,
					HargaReseller: detail.Produk.HargaReseller,
					HargaKonsumen: detail.Produk.HargaKonsumen,
					Deskripsi:     detail.Produk.Deskripsi,
					Toko: dto.TokoResponse{
						ID:       detail.Produk.Toko.ID,
						NamaToko: detail.Produk.Toko.NamaToko,
						UrlFoto:  detail.Produk.Toko.UrlFoto,
					},
					Category: dto.CategoryResponse{
						ID:           detail.Produk.Category.ID,
						NamaCategory: detail.Produk.Category.NamaCategory,
					},
					Photos: photos,
				},
				Toko: dto.TokoResponse{
					ID:       detail.Produk.Toko.ID,
					NamaToko: detail.Produk.Toko.NamaToko,
					UrlFoto:  detail.Produk.Toko.UrlFoto,
				},
				Kuantitas:  detail.Kuantitas,
				HargaTotal: detail.HargaTotal,
			}

			response.DetailTrx = append(response.DetailTrx, detailResponse)
		}

		responses = append(responses, response)
	}

	return &dto.TrxListPaginated{
		Data:  responses,
		Page:  page,
		Limit: limit,
	}, nil
}


func (s *trxService) GetTrxByID(ctx context.Context, id int64) (*dto.TrxResponse, error) {
	trx, err := s.trxRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if trx == nil {
		return nil, errors.New("trx not found")
	}

	// Mapping Entity -> DTO
	response := dto.TrxResponse{
		ID:          trx.ID,
		HargaTotal:  trx.HargaTotal,
		KodeInvoice: trx.KodeInvoice,
		MethodBayar: trx.MethodBayar,
		AlamatKirim: dto.AlamatResponse{
			ID:            trx.Alamat.ID,
			JudulAlamat:   trx.Alamat.JudulAlamat,
			NamaPenerima:  trx.Alamat.NamaPenerima,
			NoTelp:        trx.Alamat.NoTelp,
			DetailAlamat:  trx.Alamat.DetailAlamat,
		},
	}

	// Map DetailTrx
	for _, detail := range trx.Detail {
		// Map Photos
		photos := []dto.PhotoResponse{}
		for _, photo := range detail.Produk.FotoProduk {
			photos = append(photos, dto.PhotoResponse{
				ID:        photo.ID,
				ProductID: photo.IdProduk,
				Url:       photo.Url,
			})
		}

		// Map DetailTrxResponse
		detailResponse := dto.DetailTrxResponse{
			Product: dto.ProdukResponse{
				ID:            detail.Produk.ID,
				NamaProduk:    detail.Produk.NamaProduk,
				Slug:          detail.Produk.Slug,
				HargaReseller: detail.Produk.HargaReseller,
				HargaKonsumen: detail.Produk.HargaKonsumen,
				Deskripsi:     detail.Produk.Deskripsi,
				Toko: dto.TokoResponse{
					ID:       detail.Produk.Toko.ID,
					NamaToko: detail.Produk.Toko.NamaToko,
					UrlFoto:  detail.Produk.Toko.UrlFoto,
				},
				Category: dto.CategoryResponse{
					ID:           detail.Produk.Category.ID,
					NamaCategory: detail.Produk.Category.NamaCategory,
				},
				Photos: photos,
			},
			Toko: dto.TokoResponse{
				ID:       detail.Produk.Toko.ID,
				NamaToko: detail.Produk.Toko.NamaToko,
				UrlFoto:  detail.Produk.Toko.UrlFoto,
			},
			Kuantitas:  detail.Kuantitas,
			HargaTotal: detail.HargaTotal,
		}

		response.DetailTrx = append(response.DetailTrx, detailResponse)
	}

	return &response, nil
}
