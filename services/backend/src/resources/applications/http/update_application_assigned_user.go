package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/gofiber/fiber/v3"
)

func (h *ApplicationsHandler) UpdateApplicationAssignedUser(c fiber.Ctx) error {

	applicationIDString := c.Params("applicationID", "")
	if applicationIDString == "" || !uuidvalidator.ValidateUUID(applicationIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	userIDString := c.Params("userID", "")
	if userIDString != "null" && userIDString != "" && !uuidvalidator.ValidateUUID(userIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	// define empty string as null
	if userIDString == "null" || userIDString == "" {
		userIDString = ""
	}

	changeErr := h.service.UpdateApplicationAssignedUser(nil, applicationIDString, userIDString)
	if changeErr != nil {
		status, message := changeErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.SendStatus(http.StatusOK)
}
