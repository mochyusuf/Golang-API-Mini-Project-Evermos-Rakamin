package handler

import (
	"evermos_rakamin/internal/common"
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/service"
	"evermos_rakamin/internal/util"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AlamatHandler struct {
	alamatService service.AlamatService
}

func NewAlamatHandler(alamatService service.AlamatService) *AlamatHandler {
	return &AlamatHandler{alamatService}
}


func (h *AlamatHandler) GetAlamatUser(c *fiber.Ctx) error {
	

    // Type assertion with safety check
    userToken, err := common.GetUserFromToken(c)
	if err != nil {
        return util.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid token format", nil)
    }
	userid := userToken.ID
	// Fetch user profile
	alamat, err := h.alamatService.GetAlamatByUserID(c.Context(), userid)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to get user profile", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Alamat User retrieved successfully", nil, alamat)
}


func (h *AlamatHandler) GetAlamatById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid ID", err.Error(), nil)
	}

	alamat, err := h.alamatService.GetAlamatById(c.Context(), int64(id))
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to get alamat", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Alamat retrieved successfully", nil, alamat)
}

func (h *AlamatHandler) CreateAlamat(c *fiber.Ctx) error {
	var request dto.CreateAlamatRequest

	if err := c.BodyParser(&request); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}

	// Type assertion with safety check
	userToken, err := common.GetUserFromToken(c)
	if err != nil {
		return util.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid token format", nil)
	}
	request.IdUser = userToken.ID

	alamat, err := h.alamatService.CreateAlamat(c.Context(), &request)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to create alamat", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusCreated, "Alamat created successfully", nil, alamat)
}

func (h *AlamatHandler) UpdateAlamat(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid ID", err.Error(), nil)
	}
	var request dto.UpdateAlamatRequest

	if err := c.BodyParser(&request); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}


	request.ID = int64(id)

	if err := h.alamatService.UpdateAlamat(c.Context(), &request); err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to update alamat", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Alamat updated successfully", nil, nil)
}

func (h *AlamatHandler) DeleteAlamat(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid ID", err.Error(), nil)
	}

	if err := h.alamatService.DeleteAlamat(c.Context(), int64(id)); err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to delete alamat", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Alamat deleted successfully", nil, nil)
}
