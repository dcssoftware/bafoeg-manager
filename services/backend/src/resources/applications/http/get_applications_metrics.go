package http

import (
	"net/http"

	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	"github.com/gofiber/fiber/v3"
)

//	@Summary		V1 Get Applications Metrics
//	@Description	Retrieves V1 Application Metrics
//	@Tags			applications
//	@Produce		json
//	@Param			userID	query		string	false	"Filter applications by user ID"
//
// Success 		200 {object} 	models.ApplicationsMetrics "Applications Metrics"
//
//	@Failure		404		{string}	string	"The requested data could not be found."
//	@Failure		500		{string}	string	"An error occurred while processing your request. Please try again later. "
//
//	@Router			/api/v1/applications/metrics [get]
func (h *ApplicationsHandler) GetApplicationsMetrics(c fiber.Ctx) error {
	queries := c.Queries()
	userIDString := queries["userID"]

	metrics, metricsErr := h.service.GetApplicationsMetrics(nil, userIDString)
	if metricsErr != nil {
		httpStatus, userMessage := metricsErr.HTTPError()
		return c.Status(httpStatus).SendString(userMessage)
	}

	httpModel := models.ToApplicationsMetricsHttpModel(metrics)
	return c.Status(http.StatusOK).JSON(httpModel)
}
