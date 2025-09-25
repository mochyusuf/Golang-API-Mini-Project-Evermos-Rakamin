package util

import (
	"github.com/gofiber/fiber/v2"
)

type TodoResponse struct {
	Status  int        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// custom json response function from client request
func JSONResponse(c *fiber.Ctx, status int, message string, errors interface{}, data interface{}) error {
	return c.Status(status).JSON(TodoResponse{
		Status:  status,
		Message: message,
		Errors:  errors,
		Data:    data,
	})
}
