package customhttphandler

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func NotFoundHandler(c fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"code":    http.StatusNotFound,
		"message": "Not found ❌",
	})
}
