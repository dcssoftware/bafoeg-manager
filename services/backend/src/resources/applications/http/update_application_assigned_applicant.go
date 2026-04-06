package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/gofiber/fiber/v3"
)

func (h *ApplicationsHandler) UpdateApplicationAssignedApplicant(c fiber.Ctx) error {

	applicationIDString := c.Params("applicationID", "")
	applicationID, applicationIDErr := uuidvalidator.ParseHttpParamUUID(applicationIDString)
	if applicationIDString == "" || applicationIDErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	applicantIDString := c.Params("applicantID", "")
	applicantID, applicantIDErr := uuidvalidator.ParseHttpParamUUID(applicantIDString)
	if applicantIDString == "" || applicantIDErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	_, applicantModelErr := h.service.UpdateApplicationAssignedApplicant(nil, applicationID, applicantID)
	if applicantModelErr != nil {
		status, message := applicantModelErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.Status(http.StatusOK).JSON("{}")
}
