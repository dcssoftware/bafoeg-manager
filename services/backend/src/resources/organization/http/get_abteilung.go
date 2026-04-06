package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/dcssoftware/bafoeg-manager/src/resources/organization/http/models"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *OrganizationHandler) GetAbteilungen(c fiber.Ctx) error {

	queries := c.Queries()
	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(queries)
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	behördeIDString := c.Params("behoerdeID", "")
	if behördeIDString == "" || !uuidvalidator.ValidateUUID(behördeIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}
	behördeID := uuid.MustParse(behördeIDString)

	abteilungenCount, abteilungen, abteilungenErr := h.service.GetAbteilungen(nil, behördeID, pageNumber)
	if abteilungenErr != nil {
		status, message := abteilungenErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.JSON(models.ToAbteilungenResponseModel(abteilungenCount, abteilungen))
}
