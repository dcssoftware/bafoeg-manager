package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	"github.com/dcssoftware/bafoeg-manager/src/resources/organization/http/models"
	"github.com/gofiber/fiber/v3"
)

func (h *OrganizationHandler) GetRegions(c fiber.Ctx) error {

	queries := c.Queries()
	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(queries)
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	regionsCount, regions, regionsErr := h.service.GetRegions(nil, pageNumber)
	if regionsErr != nil {
		status, message := regionsErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.JSON(models.ToRegionResponseModel(regionsCount, regions))
}
