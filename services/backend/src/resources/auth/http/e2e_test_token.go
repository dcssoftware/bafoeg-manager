package http

import (
	"net/http"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/random"
	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/http/cookies"
	"github.com/gofiber/fiber/v3"
)

func (h *AuthHandler) E2ETestToken(c fiber.Ctx) error {
	jwtToken, refreshToken, err := h.service.CallbackE2ETestToken()

	if err != nil || jwtToken == "" {
		code, message := err.HTTPError()
		return c.Status(code).SendString(message)
	}

	c.Cookie(cookies.GenerateJwtToken(jwtToken, false))
	c.Cookie(cookies.GenerateRefreshToken(refreshToken, false))

	redirectURL, redirectURLErr := random.GenerateRandomHashURL(configuration.Webserver.Displayname)
	if redirectURLErr != nil {
		status, usermessage := redirectURLErr.HTTPError()
		return c.Status(status).SendString(usermessage)
	}

	return c.Redirect().Status(http.StatusMovedPermanently).To(redirectURL)
}
