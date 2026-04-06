package http

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func (h *GeneralHandler) GetApplicationInformation(c fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
