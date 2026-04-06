package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/gofiber/fiber/v3"
)

func (h *ApplicationsHandler) UpdateApplicationAssignedSchoolDegree(c fiber.Ctx) error {

	applicationIDString := c.Params("applicationID", "")
	applicationID, applicationIDErr := uuidvalidator.ParseHttpParamUUID(applicationIDString)
	if applicationIDString == "" || applicationIDErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	schoolDegreeIDString := c.Params("schoolDegreeID", "")
	schoolDegreeID, schoolDegreeIDErr := uuidvalidator.ParseHttpParamUUID(schoolDegreeIDString)
	if schoolDegreeIDString == "" || schoolDegreeIDErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	_, degreeModelErr := h.service.UpdateApplicationAssignedSchoolDegree(nil, applicationID, schoolDegreeID)
	if degreeModelErr != nil {
		status, message := degreeModelErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.Status(http.StatusOK).JSON("{}")
}
