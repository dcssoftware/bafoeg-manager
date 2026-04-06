package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	sessionlocals "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/http/consts/session-locals"
	"github.com/gofiber/fiber/v3"
)

func (h *UserHandler) GetSelfProfilePicture(c fiber.Ctx) error {
	userID, ok := c.Locals(sessionlocals.UserUUID).(string)
	if !ok {
		return c.Status(http.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
	}

	picture, pictureErr := h.service.GetProfilePictureByID(nil, userID)
	if pictureErr != nil {
		status, message := pictureErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	c.Set("Content-Type", "image/jpeg")
	return c.Status(http.StatusOK).Send(picture)
}
