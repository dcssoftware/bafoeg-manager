package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"

	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/http/cookies"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	jwt "github.com/dcssoftware/bafoeg-manager/src/helper/jwt"
	sessionLocals "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/http/consts/session-locals"
)

func (h *MiddlewareHandler) Authentication() func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		jwtToken := c.Cookies(cookies.CookieNameJwtToken, "")

		authorizationHeader := c.Get("Authorization")

		if jwtToken == "" && len(authorizationHeader) > 7 && strings.HasPrefix(authorizationHeader, "Bearer ") {
			jwtToken = authorizationHeader[7:]
		} else if jwtToken == "" {
			return c.Status(http.StatusUnauthorized).SendString(customerrorconst.NOT_AUTHORIZED_ERROR_MESSAGE)
		}

		jwtData, jwtDataErr := jwt.VerifyJWT(jwtToken)
		if jwtDataErr != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
		}

		sessionData, sessionDataErr := h.service.GetUserSessionByID(jwtData.SessionID)
		if sessionDataErr != nil || sessionData == nil || sessionData.Created.Before(time.Now().Add(-time.Hour*24*7)) {
			status, message := sessionDataErr.HTTPError()
			return c.Status(status).SendString(message)
		}

		permissions, permissionErr := h.service.GetUserPermissionsByID(jwtData.UserUUID)
		if permissionErr != nil {
			status, message := permissionErr.HTTPError()
			return c.Status(status).SendString(message)
		}

		// after jwt got evaluated, add information from jwt to fiber request context
		c.Locals(sessionLocals.UserUUID, jwtData.UserUUID)
		c.Locals(sessionLocals.Permissions, permissions)
		c.Locals(sessionLocals.JWTtoken, jwtToken)

		return c.Next()
	}
}

func arePermissionsSatisfied(wantPermissions, havePermissions []string) bool {
	foundPermission := 0

	for _, wantPermission := range wantPermissions {
	HaveLoop:
		for _, havePermission := range havePermissions {
			if wantPermission == havePermission {
				foundPermission++

				break HaveLoop
			}
		}
	}

	return foundPermission == len(wantPermissions)
}
