package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/http/models"
	"github.com/gofiber/fiber/v3"
)

// @Summary		V1 Get an Applicant (E-Akten-User) by ID
// @Description	Retrieves an V1 Applicant (E-Akten-User) for BAföG Applications
// @Tags			applications / applicants
// @Param			id	path	string	true	"Application ID"
// @Produce		json
//
// @Router			/api/v1/applications/applicants/{id} [get]
func (h *ApplicantHandler) GetApplicant(c fiber.Ctx) error {
	applicantIDString := c.Params("id", "")
	if applicantIDString == "" || !uuidvalidator.ValidateUUID(applicantIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}
	applicant, applicantErr := h.service.GetApplicantByID(nil, applicantIDString)
	if applicantErr != nil {
		status, message := applicantErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	httpModel := models.ToHttpApplicantModel(applicant)
	return c.Status(http.StatusOK).JSON(httpModel)
}
