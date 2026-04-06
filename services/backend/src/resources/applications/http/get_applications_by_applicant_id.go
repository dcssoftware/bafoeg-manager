package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	"github.com/gofiber/fiber/v3"
)

//	@Summary		V1 Get an Applications
//	@Description	Retrieves V1 Application Requests for BAföG
//	@Tags			applications
//	@Produce		json
//	@Param			page		query		int		true	"Page number for pagination"
//	@Param			applicantID	path		string	true	"Applicant ID"
//
// Success 		200 {object} 	models.ApplicationModel "Application Models"
//
//	@Failure		404			{string}	string	"The requested data could not be found."
//	@Failure		500			{string}	string	"An error occurred while processing your request. Please try again later. "
//
//	@Router			/api/v1/applications/by-applicant-id/{applicantID} [get]
func (h *ApplicationsHandler) GetApplicationsByApplicantID(c fiber.Ctx) error {
	applicantIDString := c.Params("applicantID", "")
	if applicantIDString == "" || !uuidvalidator.ValidateUUID(applicantIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(c.Queries())
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	applicationModel, applicationErr := h.service.GetApplicationsByApplicantID(nil, pageNumber, applicantIDString)
	if applicationErr != nil {
		status, message := applicationErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	applicationsCount, applicationsCountErr := h.service.GetApplicationsByApplicantIDCount(nil, applicantIDString)
	if applicationsCountErr != nil {
		status, message := applicationsCountErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	httpApplicationModel := models.ToShortHttpModels(applicationModel, applicationsCount)
	return c.
		Status(http.StatusOK).
		JSON(httpApplicationModel)
}
