package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/gofiber/fiber/v3"
)

func (h *ApplicationsHandler) GetApplicationFileByFileID(c fiber.Ctx) error {
	applicationIDString := c.Params("applicationID", "")
	if applicationIDString == "" || !uuidvalidator.ValidateUUID(applicationIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	fileIDString := c.Params("fileID", "")
	if applicationIDString == "" || !uuidvalidator.ValidateUUID(applicationIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	fileContent, fileModel, fileContentErr := h.service.GetApplicationFileByFileID(nil, fileIDString)
	if fileContentErr != nil {
		status, message := fileContentErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	if applicationIDString != fileModel.ApplicationID.String() {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	c.Set("Content-Type", fileModel.File.Type)
	return c.Status(http.StatusOK).Send(fileContent)
}
