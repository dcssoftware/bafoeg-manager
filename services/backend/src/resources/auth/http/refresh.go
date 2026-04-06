package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/http/cookies"
	sessionlocals "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/http/consts/session-locals"
	"github.com/gofiber/fiber/v3"
)

// @Summary		Get new access token via refresh token
// @Description	Reads the session out of the refresh token and recreates a new temporary access token out of it
//
// @Tags			auth
// @Produce		plain
//
// @Success		201	{string}	string
// @Failure		404	{string}	string	"The requested data could not be found."
// @Failure		500	{string}	string	"An error occurred while processing your request. Please try again later. "
//
// @Router			/api/v1/auth/refresh [get]
func (h *AuthHandler) Refresh(c fiber.Ctx) error {
	sessionID, ok := c.Locals(sessionlocals.SessionUUID).(string)
	if !ok {
		return c.Status(http.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
	}

	jwtToken, refreshTokenErr := h.service.RecreateJwtFromSession(sessionID)
	if refreshTokenErr != nil {
		httpStatus, userMessage := refreshTokenErr.HTTPError()
		return c.Status(httpStatus).SendString(userMessage)
	}

	c.Cookie(cookies.GenerateJwtToken(jwtToken, false))
	return c.Status(http.StatusCreated).SendString(jwtToken)
}
