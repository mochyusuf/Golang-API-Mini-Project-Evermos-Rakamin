package handler

import (
	"evermos_rakamin/internal/util"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProxyHandler struct {
}

func NewProxyHandler() *ProxyHandler {
	return &ProxyHandler{}
}

func (h *ProxyHandler) GetAllProvinces(c *fiber.Ctx) error {
	provinces, err := util.GetAllProvinces()
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to fetch provinces", err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Succeed to GET data", nil, provinces)
}

func (h *ProxyHandler) GetCitiesByProvinceID(c *fiber.Ctx) error {
	provinceID := c.Params("prov_id")
	if provinceID == "" {
		return util.JSONResponse(c, http.StatusBadRequest, "Province ID is required", "Invalid Province ID", nil)
	}

	cities, err := util.GetAllCitiesByProvinceID(provinceID)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to fetch cities", err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Succeed to GET data", nil, cities)
}

func (h *ProxyHandler) GetProvinceByID(c *fiber.Ctx) error {
	id := c.Params("prov_id")
	if id == "" {
		return util.JSONResponse(c, http.StatusBadRequest, "Province ID is required", "Invalid Province ID", nil)
	}

	province, err := util.GetProvinceByID(id)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to fetch province", err.Error(), nil)
	}
	if province == nil {
		return util.JSONResponse(c, http.StatusNotFound, "Province not found", "No data found for the given ID", nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Succeed to GET data", nil, province)
}

func (h *ProxyHandler) GetCityByID(c *fiber.Ctx) error {
	id := c.Params("city_id")
	provID := c.Params("prov_id")
	if id == "" && provID == "" {
		return util.JSONResponse(c, http.StatusBadRequest, "City ID is required", "Invalid City ID", nil)
	}

	city, err := util.GetCityByID(provID, id)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to fetch city", err.Error(), nil)
	}
	if city == nil {
		return util.JSONResponse(c, http.StatusNotFound, "City not found", "No data found for the given ID", nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Succeed to GET data", nil, city)
}

func (h *ProxyHandler) GetCityByIDOnly(c *fiber.Ctx) error {
	id := c.Params("city_id")
	if id == "" {
		return util.JSONResponse(c, http.StatusBadRequest, "City ID is required", "Invalid City ID", nil)
	}

	city, err := util.GetCityByIDOnly(id)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to fetch city", err.Error(), nil)
	}
	if city == nil {
		return util.JSONResponse(c, http.StatusNotFound, "City not found", "No data found for the given ID", nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Succeed to GET data", nil, city)
}