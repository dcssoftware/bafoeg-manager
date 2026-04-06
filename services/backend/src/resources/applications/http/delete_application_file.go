package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/gofiber/fiber/v3"
)

func (h *ApplicationsHandler) DeleteApplicationFile(c fiber.Ctx) error {
	applicationIDString := c.Params("applicationID", "")
	if applicationIDString == "" || !uuidvalidator.ValidateUUID(applicationIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	fileIDString := c.Params("fileID", "")
	if fileIDString == "" || !uuidvalidator.ValidateUUID(fileIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	err := h.service.DeleteApplicationFile(nil, applicationIDString, fileIDString)
	if err != nil {
		return nil
	}

	return nil
}
