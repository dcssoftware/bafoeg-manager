package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	"github.com/dcssoftware/bafoeg-manager/src/resources/application-labels/http/models"
	"github.com/gofiber/fiber/v3"
)

func (h *ApplicationLabelsHandler) GetApplicationLabels(c fiber.Ctx) error {

	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(c.Queries())
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	labels, labelsErr := h.service.GetLabels(nil, pageNumber)
	if labelsErr != nil {
		status, message := labelsErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	httpLabels := models.ToHttpModels(labels)

	return c.Status(http.StatusOK).JSON(httpLabels)
}
