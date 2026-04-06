package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/http/models"
	"github.com/gofiber/fiber/v3"
)

//	@Summary		V1 Get schools
//	@Description	Retrieves V1 Schools
//	@Tags			schools
//	@Produce		json
//	@Param			page	query		int		true	"Page number for pagination"
//
// Success 		200 {object} 	models.SchoolShortResponseModel "School Model"
//
//	@Failure		404		{string}	string	"The requested data could not be found."
//	@Failure		500		{string}	string	"An error occurred while processing your request. Please try again later. "
//
//	@Router			/api/v1/schools/ [get]
func (h *SchoolHandler) GetSchools(c fiber.Ctx) error {

	queries := c.Queries()
	filterResultString := queries["filterResult"]
	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(c.Queries())
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	schools, schoolsErr := h.service.GetSchools(nil, pageNumber, filterResultString)
	if schoolsErr != nil {
		status, message := schoolsErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	schoolsCount, schoolsCountErr := h.service.GetSchoolsCount(nil, filterResultString)
	if schoolsCountErr != nil {
		status, message := schoolsCountErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	httpSchools := models.ToHttpSchoolShortModels(schools, schoolsCount)
	return c.Status(http.StatusOK).JSON(httpSchools)
}
