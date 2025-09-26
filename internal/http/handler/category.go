package handler

import (
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/service"
	"evermos_rakamin/internal/util"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService}
}

func (h *CategoryHandler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := h.categoryService.GetAllCategories(c.Context())
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to fetch categories", err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Success", nil, categories)
}

func (h *CategoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid ID", err.Error(), nil)
	}

	category, err := h.categoryService.GetCategoryByID(c.Context(), id)
	if err != nil {
		return util.JSONResponse(c, http.StatusNotFound, "Category not found", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Success", nil, category)
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var req dto.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}

	if err := h.categoryService.CreateCategory(c.Context(), &req); err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to create category", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusCreated, "Category created successfully", nil, nil)
}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid ID", err.Error(), nil)
	}

	var req dto.UpdateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}

	if err := h.categoryService.UpdateCategory(c.Context(), id, &req); err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to update category", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Category updated successfully", nil, nil)
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid ID", err.Error(), nil)
	}

	if err := h.categoryService.DeleteCategory(c.Context(), id); err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to delete category", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Category deleted successfully", nil, nil)
}