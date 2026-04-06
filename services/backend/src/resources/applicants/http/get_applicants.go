package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/http/models"
	"github.com/gofiber/fiber/v3"
)

// @Summary		V1 Get Applicants (E-Akten-Users)
// @Description	Retrieves V1 Applicants (E-Akten-Users)for BAföG Applications
// @Tags			applications / applicants
// @Produce		json
//
// @Param			page	query	int	true	"Page number for pagination"
//
// @Router			/api/v1/applications/applicants [get]
func (h *ApplicantHandler) GetApplicants(c fiber.Ctx) error {
	queries := c.Queries()
	filterResultString := queries["filterResult"]

	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(queries)
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	applicant, applicantErr := h.service.GetApplicants(nil, pageNumber, filterResultString)
	if applicantErr != nil {
		status, message := applicantErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	applicantCount, applicantCountErr := h.service.GetApplicantsCount(nil, filterResultString)
	if applicantCountErr != nil {
		status, message := applicantCountErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	httpModel := models.ToHttpApplicantModels(applicant, applicantCount)
	return c.Status(http.StatusOK).JSON(httpModel)
}
