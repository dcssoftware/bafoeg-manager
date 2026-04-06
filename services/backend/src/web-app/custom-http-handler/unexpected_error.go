package customhttphandler

import "github.com/gofiber/fiber/v3"

func UnexpectedErrorHandler(c fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Unexpected Error Occurred")
}
