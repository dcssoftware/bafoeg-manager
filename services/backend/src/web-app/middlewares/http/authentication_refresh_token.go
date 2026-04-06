package middlewares

import (
	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/http/cookies"
	"github.com/gofiber/fiber/v3"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	jwt "github.com/dcssoftware/bafoeg-manager/src/helper/jwt"
	sessionLocals "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/http/consts/session-locals"
)

func (h *MiddlewareHandler) AuthenticationRefreshToken() func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		refreshToken := c.Cookies(cookies.CookieNameRefreshToken, "")

		if refreshToken == "" {
			return c.Status(fiber.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
		}

		jwtData, jwtDataErr := jwt.VerifyRefreshToken(refreshToken)
		if jwtDataErr != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
		}

		// after jwt got evaluated, add information from jwt to fiber request context
		c.Locals(sessionLocals.SessionUUID, jwtData.SessionID)

		return c.Next()
	}
}
