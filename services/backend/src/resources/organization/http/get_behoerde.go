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

func (h *OrganizationHandler) GetBehörden(c fiber.Ctx) error {

	queries := c.Queries()
	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(queries)
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	regionIDString := c.Params("regionID", "")
	if regionIDString == "" || !uuidvalidator.ValidateUUID(regionIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}
	regionID := uuid.MustParse(regionIDString)

	behördenCount, behörden, behördenErr := h.service.GetBehörden(nil, regionID, pageNumber)
	if behördenErr != nil {
		status, message := behördenErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.JSON(models.ToBehördeResponseModel(behördenCount, behörden))
}
