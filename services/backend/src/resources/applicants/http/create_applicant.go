package http

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	custombadrequestconstraints "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/bad-request-constraints"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/gofiber/fiber/v3"
)

func (h *ApplicantHandler) CreateApplicant(c fiber.Ctx) error {

	createApplicantModel := new(models.CreateApplicantModel)
	bodyParseErr := c.Bind().Body(createApplicantModel)
	if bodyParseErr != nil {
		status, message := customerrors.
			NewBadRequestError(custombadrequestconstraints.BadRequest_ApplicationStatusChangeInvalid).
			HTTPError()
		return c.Status(status).SendString(message)
	}

	h.service.InsertApplicant(nil, models.CreateApplicantModel{
		Firstname: createApplicantModel.Firstname,
		Lastname:  createApplicantModel.Lastname,

		Street:      createApplicantModel.Street,
		HouseNumber: createApplicantModel.HouseNumber,
		ZipCode:     createApplicantModel.ZipCode,
		City:        createApplicantModel.City,
		Country:     createApplicantModel.Country,
	})

	return nil
}
