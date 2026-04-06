package http

import (
	"github.com/dcssoftware/bafoeg-manager/src/resources/user/http/models"
	"github.com/gofiber/fiber/v3"
)

func (h *UserHandler) GetUsers(c fiber.Ctx) error {

	userModels, userModelErr := h.service.GetUsers(nil)
	if userModelErr != nil {
		status, message := userModelErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	httpUserModels := models.ToHttpUserModels(userModels)
	return c.JSON(httpUserModels)
}
