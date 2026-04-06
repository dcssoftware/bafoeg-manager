package http

import (
	"net/http"
	"strconv"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/http/models"
	"github.com/gofiber/fiber/v3"
)

// @Summary		V1 Get an Applicants (E-Akten-User) by School ID
// @Description	Retrieves an V1 Applicants (E-Akten-User) by School ID
// @Tags			applications / applicants
// @Param			schoolID	path	string	true	"School ID"
// @Param			page		query	int		true	"Page number for pagination"
// @Param			isActive	query	int		true	"Show only active students (applicants)"
//
// @Produce		json
//
// @Router			/api/v1/applications/applicants/by-school/{schoolID} [get]
func (h *ApplicantHandler) GetApplicantsBySchoolID(c fiber.Ctx) error {
	queries := c.Queries()
	schoolIDString := c.Params("schoolID", "")

	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(queries)
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	isActiveString := queries["isActive"]
	isActive, isActiveErr := strconv.ParseBool(isActiveString)
	if isActiveErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	if schoolIDString == "" || !uuidvalidator.ValidateUUID(schoolIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	applicants, applicantsErr := h.service.GetApplicantsBySchoolID(nil, pageNumber, isActive, schoolIDString)
	if applicantsErr != nil {
		httpStatus, userMessage := applicantsErr.HTTPError()
		return c.Status(httpStatus).SendString(userMessage)
	}

	applicantsCount, applicantsCountErr := h.service.GetApplicantsBySchoolIDCount(nil, isActive, schoolIDString)
	if applicantsCountErr != nil {
		httpStatus, userMessage := applicantsCountErr.HTTPError()
		return c.Status(httpStatus).SendString(userMessage)
	}

	httpModel := models.ToHttpApplicantsBySchoolModels(applicants, applicantsCount)
	return c.Status(http.StatusOK).JSON(httpModel)
}
