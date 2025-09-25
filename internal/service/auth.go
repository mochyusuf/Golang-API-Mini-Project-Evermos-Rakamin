package service

import (
	"evermos_rakamin/internal/config"
	"evermos_rakamin/internal/common"
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/entity"
	"evermos_rakamin/internal/repository"
	"evermos_rakamin/internal/util"
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error)
	Register(ctx context.Context, request dto.RegisterRequest) (string, error)
}

type authService struct {
	cfg        *config.Config
	repository repository.UserRepository
}

func NewAuthService(cfg *config.Config, repository repository.UserRepository) AuthService {
	return &authService{cfg, repository}
}

func (u *authService) Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.repository.FindByContact(ctx, request.NoTelp)

	if err != nil {
		return nil, errors.New("No Telepon/kata sandi salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(request.KataSandi))

	if err != nil {
		return nil, errors.New("No Telepon/kata sandi salah")
	}

	token, err := common.GenerateAccessToken(ctx, user)
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


	response := &dto.LoginResponse{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: user.TanggalLahir.Format("01/01/2001"),
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
		Token: token,
	}

	return response, nil
}

func (u *authService) Register(ctx context.Context, request dto.RegisterRequest) (string, error) {
	tanggalLahir, err := time.Parse("2001-01-01", request.TanggalLahir)
	isAdmin := false

	// Lengkapi data user dari request
	user := &entity.User{
		Nama:         request.Nama,
		Email:        request.Email,
		KataSandi:    request.KataSandi,
		NoTelp:       request.NoTelp,
		IsAdmin:      isAdmin,
		TanggalLahir: tanggalLahir,
		JenisKelamin: request.JenisKelamin,
		Tentang:      request.Tentang,
		Pekerjaan:    request.Pekerjaan,
		IdProvinsi:   request.IdProvinsi,
		IdKota:       request.IdKota,
	}

	

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.KataSandi = string(hashedPassword)

	if err := u.repository.Create(ctx, user); err != nil {
		return "", err
	}

	token, err := common.GenerateAccessToken(ctx, user)
	if err != nil {
		return "", err
	}

	return token, nil
}