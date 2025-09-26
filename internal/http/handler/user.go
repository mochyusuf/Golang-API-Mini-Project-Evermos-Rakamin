package handler

import (
	"evermos_rakamin/internal/common"
	"evermos_rakamin/internal/dto"
	"evermos_rakamin/internal/service"
	"evermos_rakamin/internal/util"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GeneratePassword(c *fiber.Ctx) error {
	var request struct {
		Password string `json:"kata_sandi"`
	}

	if err := c.BodyParser(&request); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to generate password", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Successfully generated password", nil, string(hashedPassword))
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {

    // Type assertion with safety check
    userToken, err := common.GetUserFromToken(c)
	if err != nil {
        return util.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid token format", nil)
    }

	userContact := userToken.NoTelp

	// Fetch user profile
	user, err := h.userService.GetProfile(c.Context(), userContact)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to get user profile", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "User profile retrieved successfully", nil, user)
}

func (h *UserHandler) GetAllUser(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers(c.Context())
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to get users", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Users retrieved successfully", nil, users)
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	var request dto.UpdateUserRequest

	if err := c.BodyParser(&request); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}

	// Type assertion with safety check
	userToken, err := common.GetUserFromToken(c)
	if err != nil {
		return util.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid token format", nil)
	}
	userid := userToken.ID

	err = h.userService.UpdateProfile(c.Context(), userid, &request)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to update profile", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Profile updated successfully", nil, nil)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var request dto.UpdateUserRequest

	if err := c.BodyParser(&request); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}

	err = h.userService.UpdateProfile(c.Context(), int64(id), &request)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, "Failed to update user", err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "User updated successfully", nil, nil)
}