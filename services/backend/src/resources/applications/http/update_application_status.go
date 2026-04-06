package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	applicationStates "github.com/dcssoftware/bafoeg-manager/src/resources/applications/states"
	"github.com/gofiber/fiber/v3"
)

func (h *ApplicationsHandler) UpdateApplicationStatus(c fiber.Ctx) error {

	applicationIDString := c.Params("applicationID", "")
	if applicationIDString == "" || !uuidvalidator.ValidateUUID(applicationIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	statusString := c.Params("status", "")
	newApplicationState, newApplicationStateErr := applicationStates.ConvertStrToApplicationState(statusString)
	if newApplicationStateErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	changeErr := h.service.UpdateApplicationStatus(nil, applicationIDString, newApplicationState)
	if changeErr != nil {
		status, message := changeErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.SendStatus(http.StatusOK)
}
