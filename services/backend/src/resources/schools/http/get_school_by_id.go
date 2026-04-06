package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/http/models"
	"github.com/gofiber/fiber/v3"
)

//	@Summary		V1 Get school by ID
//	@Description	Retrieves V1 School by ID
//	@Tags			schools
//	@Produce		json
//	@Param			schoolID	path		string	true	"School ID"
//
// Success 		200 {object} 	models.SchoolModel "School Model"
//
//	@Failure		404			{string}	string	"The requested data could not be found."
//	@Failure		500			{string}	string	"An error occurred while processing your request. Please try again later. "
//
//	@Router			/api/v1/schools/{schoolID} [get]
func (h *SchoolHandler) GetSchoolByID(c fiber.Ctx) error {

	schoolIDString := c.Params("schoolID", "")
	if schoolIDString == "" || !uuidvalidator.ValidateUUID(schoolIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	schools, schoolsErr := h.service.GetSchoolByID(schoolIDString)
	if schoolsErr != nil {
		status, message := schoolsErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	httpSchools := models.ToHttpSchoolModel(schools)
	return c.Status(http.StatusOK).JSON(httpSchools)
}
