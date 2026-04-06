package http

import (
	"net/http"

	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/http/cookies"
	"github.com/gofiber/fiber/v3"
)

func (h *AuthHandler) Logout(c fiber.Ctx) error {
	c.Cookie(cookies.GenerateRefreshToken("", true))
	c.Cookie(cookies.GenerateJwtToken("", true))

	return c.Redirect().Status(http.StatusMovedPermanently).To("/")
}
