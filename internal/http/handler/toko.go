package handler

import (
	"evermos_rakamin/internal/common"
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/service"
	"evermos_rakamin/internal/util"
	"strconv"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TokoHandler struct {
	tokoService service.TokoService
}

func NewTokoHandler(tokoService service.TokoService) *TokoHandler {
	return &TokoHandler{tokoService}
}

func (h *TokoHandler) GetMyToko(c *fiber.Ctx) error {
	// Type assertion with safety check
	userToken, err := common.GetUserFromToken(c)
	if err != nil {
		return util.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid token format", nil)
	}

	// Fetch toko by user ID
	toko, err := h.tokoService.GetTokoByUserID(c.Context(), userToken.ID)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to get toko", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Toko retrieved successfully", nil, toko)
}

func (h *TokoHandler) GetTokoByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id_toko")
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid toko ID", err.Error(), nil)
	}

	toko, err := h.tokoService.GetTokoByID(c.Context(), int64(id))
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to get toko", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Toko retrieved successfully", nil, toko)
}

func (h *TokoHandler) GetTokoPaginated(c *fiber.Ctx) error {
	// Parse query parameters for pagination

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	nama := c.Query("nama", "")

	result, err := h.tokoService.GetTokoPaginated(c.Context(), limit, page, nama)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to get paginated toko", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Toko paginated retrieved successfully", nil, result)
}

func (h *TokoHandler) UpdateToko(c *fiber.Ctx) error {
	var request dto.UpdateTokoRequest

	if err := c.BodyParser(&request); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}

	id, err := c.ParamsInt("id_toko")
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid toko ID", err.Error(), nil)
	}

	// Ambil file dari form-data
	fileHeader, err := c.FormFile("photo")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to get file"})
	}

	// Open file untuk mendapatkan io.Reader
	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to open file"})
	}
	defer file.Close()

	// Upload
	urlFoto, err := util.UploadFileToko(c, fileHeader, fileHeader.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to upload file", "error": err.Error()})
	}

	// Update URL Foto ke Database
	request.UrlFoto = &urlFoto
	err = h.tokoService.UpdateToko(c.Context(), int64(id), &request)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to update toko", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Toko updated successfully", nil, urlFoto)
}
