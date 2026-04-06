package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/resources/user/http/models"
	sessionlocals "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/http/consts/session-locals"
	"github.com/gofiber/fiber/v3"
)

// @Summary		Get own user permissions
// @Description	Returns all set permission for authorized users
//
// @Tags			user
// @Produce		json
//
// @Success		200	{object}	models.SelfInformationModel
// @Failure		404	{string}	string	"The requested data could not be found."
// @Failure		500	{string}	string	"An error occurred while processing your request. Please try again later. "
//
// @Router			/api/v1/self [get]
func (s *UserHandler) GetSelfInformation(c fiber.Ctx) error {
	userID, ok := c.Locals(sessionlocals.UserUUID).(string)
	if !ok {
		return c.Status(http.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
	}

	selfInformation, selfInformationErr := s.service.GetSelfInformationByID(userID)
	if selfInformationErr != nil {
		status, message := selfInformationErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	httpModel := models.NewSelfInformationModel(selfInformation)
	return c.Status(http.StatusOK).JSON(httpModel)
}
