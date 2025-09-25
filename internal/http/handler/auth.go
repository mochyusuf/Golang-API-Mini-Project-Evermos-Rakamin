package handler

import (
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/service"
	"evermos_rakamin/internal/util"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h * AuthHandler) Login(c *fiber.Ctx) error {
	var request dto.LoginRequest

	// Correct method for parsing JSON body
	if err := c.BodyParser(&request); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	// Call the Login method of the authService
	response, err := h.authService.Login(c.Context(), request)
	if err != nil {
		return util.JSONResponse(c, http.StatusUnauthorized, "Unauthoqize",err.Error(), nil,)
	}


	return util.JSONResponse(c, http.StatusOK,"Login successful",nil, response)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var request dto.RegisterRequest

	// Correct method for parsing JSON body
	if err := c.BodyParser(&request); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}

	token, err := h.authService.Register(c.Context(), request)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Registration failed", err.Error(), nil)
	}

	data := map[string]string{"token": token}

	return util.JSONResponse(c, http.StatusCreated, "Registration successful", nil, data)
}