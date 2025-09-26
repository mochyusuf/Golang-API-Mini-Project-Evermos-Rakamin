package service

import (
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/entity"
	"evermos_rakamin/internal/repository"
	"context"
	"errors"
)

type CategoryService interface {
	GetAllCategories(ctx context.Context) ([]dto.CategoryResponse, error)
	GetCategoryByID(ctx context.Context, id int64) (*dto.CategoryResponse, error)
	CreateCategory(ctx context.Context, req *dto.CreateCategoryRequest) error
	UpdateCategory(ctx context.Context, id int64, req *dto.UpdateCategoryRequest) error
	DeleteCategory(ctx context.Context, id int64) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo}
}

func (s *categoryService) GetAllCategories(ctx context.Context) ([]dto.CategoryResponse, error) {
	categories, err := s.categoryRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.CategoryResponse, len(categories))
	for i, c := range categories {
		responses[i] = dto.CategoryResponse{
			ID:           c.ID,
			NamaCategory: c.NamaCategory,
		}
	}

	return responses, nil
}

func (s *categoryService) GetCategoryByID(ctx context.Context, id int64) (*dto.CategoryResponse, error) {
	category, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}

	return &dto.CategoryResponse{
		ID:           category.ID,
		NamaCategory: category.NamaCategory,
	}, nil
}

func (s *categoryService) CreateCategory(ctx context.Context, req *dto.CreateCategoryRequest) error {
	category := &entity.Category{
		NamaCategory: req.NamaCategory,
	}
	return s.categoryRepo.Create(ctx, category)
}

func (s *categoryService) UpdateCategory(ctx context.Context, id int64, req *dto.UpdateCategoryRequest) error {
	category, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if category == nil {
		return errors.New("category not found")
	}

	if req.NamaCategory != nil {
		category.NamaCategory = *req.NamaCategory
	}

	return s.categoryRepo.Update(ctx, category)
}

func (s *categoryService) DeleteCategory(ctx context.Context, id int64) error {
	return s.categoryRepo.Delete(ctx, id)
}