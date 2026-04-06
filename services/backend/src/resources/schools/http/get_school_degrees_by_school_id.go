package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/http/models"
	"github.com/gofiber/fiber/v3"
)

func (h *SchoolHandler) GetSchoolDegreesBySchoolID(c fiber.Ctx) error {

	schoolIDString := c.Params("schoolID", "")
	if schoolIDString == "" || !uuidvalidator.ValidateUUID(schoolIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(c.Queries())
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	degrees, degreesErr := h.service.GetSchoolDegreeBySchoolID(nil, pageNumber, schoolIDString)
	if degreesErr != nil {
		status, message := degreesErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	degreesCount, degreesCountErr := h.service.GetSchoolDegreeBySchoolIDCount(nil, schoolIDString)
	if degreesCountErr != nil {
		status, message := degreesCountErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.JSON(models.ToHttpSchoolDegreeResponseModel(degrees, degreesCount))
}
