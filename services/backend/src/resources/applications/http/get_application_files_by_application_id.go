package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	"github.com/gofiber/fiber/v3"
)

func (h *ApplicationsHandler) GetApplicationFilesByApplicationID(c fiber.Ctx) error {
	applicationIDString := c.Params("applicationID", "")
	if applicationIDString == "" || !uuidvalidator.ValidateUUID(applicationIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(c.Queries())
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	files, filesErr := h.service.GetApplicationFilesByApplicationID(nil, pageNumber, applicationIDString)
	if filesErr != nil {
		status, message := filesErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	filesCount, filesCountErr := h.service.GetApplicationFilesByApplicationIDCount(nil, applicationIDString)
	if filesCountErr != nil {
		status, message := filesCountErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	httpApplicationModel := models.ToFilesHttpModels(files, filesCount)

	return c.Status(http.StatusOK).JSON(httpApplicationModel)
}
