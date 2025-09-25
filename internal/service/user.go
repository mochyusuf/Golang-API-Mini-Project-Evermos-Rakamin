package service

import (
	"evermos_rakamin/internal/config"
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/entity"
	"evermos_rakamin/internal/repository"
	"evermos_rakamin/internal/util"
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetProfile(ctx context.Context, NoTelp string) (*dto.ProfileResponse, error)
	UpdateProfile(ctx context.Context, id int64, req *dto.UpdateUserRequest) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type userService struct {
	cfg        *config.Config
	userRepo repository.UserRepository
}

func NewUserService(cfg *config.Config,userRepo repository.UserRepository) UserService {
	return &userService{cfg, userRepo}
}

func (s *userService) GetProfile(ctx context.Context, NoTelp string) (*dto.ProfileResponse, error) {
	user, err := s.userRepo.FindByContact(ctx, NoTelp)
	if err != nil {
		return nil, err
	}

	province, err := util.GetProvinceByID(user.IdProvinsi)
	if err != nil {
		return nil, err
	}

	city, err := util.GetCityByID(user.IdProvinsi, user.IdKota)
	if err != nil {
		return nil, err
	}

	// Mapping
	profileResp := &dto.ProfileResponse{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: user.TanggalLahir.Format("02/01/2006"),
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IdProvinsi: dto.Province{
			ID:   province.ID,
			Name: province.Name,
		},
		IdKota: dto.City{
			ID:         city.ID,
			ProvinceID: city.ProvinceID,
			Name:       city.Name,
		},
	}

	return profileResp, nil
}

func (s *userService) UpdateProfile(ctx context.Context, id int64, req *dto.UpdateUserRequest) error {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	tanggalLahir, err := time.Parse("2001-01-01", *req.TanggalLahir)

	if req.Nama != nil {
		user.Nama = *req.Nama
	}
	if req.KataSandi != nil {
		user.KataSandi = *req.KataSandi
	}
	if req.NoTelp != nil {
		user.NoTelp = *req.NoTelp
	}
	if req.TanggalLahir != nil {
		user.TanggalLahir = tanggalLahir
	}
	if req.JenisKelamin != nil {
		user.JenisKelamin = *req.JenisKelamin
	}
	if req.Tentang != nil {
		user.Tentang = *req.Tentang
	}
	if req.Pekerjaan != nil {
		user.Pekerjaan = *req.Pekerjaan
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.IdProvinsi != nil {
		user.IdProvinsi = *req.IdProvinsi
	}
	if req.IdCity != nil {
		user.IdKota = *req.IdCity
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.KataSandi = string(hashedPassword)

	return s.userRepo.Update(ctx, user)
}


func (s *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return s.userRepo.FindAll(ctx)
}