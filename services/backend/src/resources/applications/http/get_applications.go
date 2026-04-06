package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	"github.com/gofiber/fiber/v3"
)

//	@Summary		V1 Get an Applications
//	@Description	Retrieves V1 Application Requests for BAföG
//	@Tags			applications
//	@Produce		json
//	@Param			page	query		int		true	"Page number for pagination"
//	@Param			userID	query		string	false	"Filter applications by user ID"
//
// Success 		200 {object} 	models.ApplicationModel "Application Models"
//
//	@Failure		404		{string}	string	"The requested data could not be found."
//	@Failure		500		{string}	string	"An error occurred while processing your request. Please try again later. "
//
//	@Router			/api/v1/applications/ [get]
func (h *ApplicationsHandler) GetApplications(c fiber.Ctx) error {
	queries := c.Queries()
	userIDString := queries["userID"]
	filterResultString := queries["filterResult"]
	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(queries)
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	var showFinishedApplicationsDefault bool = false
	showFinishedApplications, showFinishedApplicationsErr := httpParams.GetParamsBoolean(queries, "showAllApplications", &showFinishedApplicationsDefault)
	if showFinishedApplicationsErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	applicationModel, applicationErr := h.service.GetApplications(pageNumber, userIDString, !showFinishedApplications, filterResultString)
	if applicationErr != nil {
		status, message := applicationErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	applicationsCount, applicationsCountErr := h.service.GetApplicationsCount(userIDString, "")
	if applicationsCountErr != nil {
		status, message := applicationsCountErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	httpApplicationModel := models.ToShortHttpModels(applicationModel, applicationsCount)
	return c.
		Status(http.StatusOK).
		JSON(httpApplicationModel)
}
