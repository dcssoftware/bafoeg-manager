package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	"github.com/gofiber/fiber/v3"
)

//	@Summary		V1 Get an Application by ID
//	@Description	Retrieves V1 Application Request by ID for BAföG
//	@Tags			applications
//	@Produce		json
//	@Param			applicationID	path		string	true	"Application ID"
//
// Success 		200 {object} 	models.ApplicationModel "Application Model"
//
//	@Failure		404				{string}	string	"The requested data could not be found."
//	@Failure		500				{string}	string	"An error occurred while processing your request. Please try again later. "
//	@Router			/api/v1/applications/{applicationID} [get]
func (h *ApplicationsHandler) GetApplicationByID(c fiber.Ctx) error {

	applicationIDString := c.Params("applicationID", "")
	if applicationIDString == "" || !uuidvalidator.ValidateUUID(applicationIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	application, applicationErr := h.service.GetApplicationByIDWithTrainingsAddress(nil, applicationIDString)
	if applicationErr != nil {
		httpStatus, userMessage := applicationErr.HTTPError()
		return c.Status(httpStatus).SendString(userMessage)
	}

	httpApplication := models.ToHttpModel(application)
	return c.Status(http.StatusOK).JSON(httpApplication)
}
