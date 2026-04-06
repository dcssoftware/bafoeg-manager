package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	"github.com/gofiber/fiber/v3"
)

//	@Summary		V1 Get Application Revisions by Application ID
//	@Description	Retrieves V1 Application Revisions by Application ID for BAföG
//	@Tags			applications / revision
//	@Produce		json
//	@Param			page			query		int		true	"Page number for pagination"
//	@Param			applicationID	path		string	true	"Application ID"
//
// Success 		200 {object} 	models.ApplicationRevisionShortModels "Application Revision Model"
//
//	@Failure		404				{string}	string	"The requested data could not be found."
//	@Failure		500				{string}	string	"An error occurred while processing your request. Please try again later. "
//	@Router			/api/v1//applications/{applicationID}/revision [get]
func (h *ApplicationsHandler) GetApplicationRevisionsByApplicationID(c fiber.Ctx) error {
	applicationIDString := c.Params("applicationID", "")
	if applicationIDString == "" || !uuidvalidator.ValidateUUID(applicationIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(c.Queries())
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	revisionsModels, revisionsModelsErr := h.service.GetApplicationRevisionsByApplicationID(nil, pageNumber, applicationIDString)
	if revisionsModelsErr != nil {
		httpStatus, userMessage := revisionsModelsErr.HTTPError()
		return c.Status(httpStatus).SendString(userMessage)
	}

	revisionsCount, revisionsCountErr := h.service.GetApplicationRevisionsByApplicationIDCount(nil, applicationIDString)
	if revisionsCountErr != nil {
		httpStatus, userMessage := revisionsCountErr.HTTPError()
		return c.Status(httpStatus).SendString(userMessage)
	}

	httpModel := models.ToRevisionShortHttpModels(revisionsModels, revisionsCount)
	return c.Status(http.StatusOK).JSON(httpModel)
}
