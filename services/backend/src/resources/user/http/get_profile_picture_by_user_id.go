package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/gofiber/fiber/v3"
)

func (h *UserHandler) GetProfilePictureByUserID(c fiber.Ctx) error {
	userIDString := c.Params("userid", "")
	if userIDString == "" || !uuidvalidator.ValidateUUID(userIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	picture, pictureErr := h.service.GetProfilePictureByID(nil, userIDString)
	if pictureErr != nil {
		status, message := pictureErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	c.Set("Content-Type", "image/jpeg")
	return c.Status(http.StatusOK).Send(picture)
}
