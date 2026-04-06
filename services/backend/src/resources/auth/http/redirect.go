package http

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/http/cookies"
	"github.com/gofiber/fiber/v3"
)

func (h *AuthHandler) CreateRedirect(c fiber.Ctx) error {
	randomBytesLength := 16
	randomBytes := make([]byte, randomBytesLength)
	_, readErr := rand.Read(randomBytes)
	if readErr != nil {
		httpStatus, httpMessage := customerrors.NewInternalServerError(
			readErr,
			"",
			"Could not create login hash",
		).HTTPError()
		return c.Status(httpStatus).SendString(httpMessage)
	}

	state := base64.URLEncoding.EncodeToString(randomBytes)

	redirectURL, redirectURLErr := h.service.CreateRedirect(configuration.Authentication.Oauth[0].Identifier, state)
	if redirectURLErr != nil {
		status, message := redirectURLErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	c.Cookie(cookies.GenerateOauthStateCookie(state, false))
	c.Cookie(cookies.GenerateAuthProviderCookie(configuration.Authentication.Oauth[0].Identifier, false)) // Todo: make it dynamic

	return c.Redirect().Status(http.StatusTemporaryRedirect).To(redirectURL)
}
