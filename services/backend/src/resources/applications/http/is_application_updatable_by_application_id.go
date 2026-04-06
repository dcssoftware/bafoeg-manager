package http

import (
	"net/http"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/gofiber/fiber/v3"
)

func (h *ApplicationsHandler) IsApplicationUpdatableByApplicationID(c fiber.Ctx) error {

	applicationIDString := c.Params("applicationID", "")
	applicationID, applicationIDErr := uuidvalidator.ParseHttpParamUUID(applicationIDString)
	if applicationIDString == "" || applicationIDErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	isUpdatable, _, isUpdatabaleErr := h.service.IsApplicationUpdatableByApplicationID(nil, applicationID)
	if isUpdatabaleErr != nil {
		status, message := isUpdatabaleErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	if !isUpdatable {
		status, message := customerrors.NewApplicationAlreadyProcessedError().HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.Status(http.StatusOK).JSON("{}")
}
