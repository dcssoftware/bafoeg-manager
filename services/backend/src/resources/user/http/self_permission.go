package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	models "github.com/dcssoftware/bafoeg-manager/src/resources/user/http/models"
	sessionLocals "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/http/consts/session-locals"
	"github.com/gofiber/fiber/v3"
)

// @Summary		Get own user permissions
// @Description	Returns all set permission for authorized users
//
// @Tags			user
// @Produce		plain
//
// @Success		200	{object}	models.SelfPermissionsModel
// @Failure		404	{string}	string	"The requested data could not be found."
// @Failure		500	{string}	string	"An error occurred while processing your request. Please try again later. "
//
// @Router			/api/v1/self/permissions [get]
func (s *UserHandler) GetSelfPermissions(c fiber.Ctx) error {
	userID, ok := c.Locals(sessionLocals.UserUUID).(string)
	if !ok {
		return c.Status(http.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
	}

	permissions, permissionsErr := s.service.GetUserPermissionsByID(nil, userID)
	if permissionsErr != nil {
		status, message := permissionsErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	permissionModel := models.NewSelfPermissionsModel(permissions)
	return c.Status(http.StatusOK).JSON(permissionModel)
}
