package middlewares

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	sessionLocals "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/http/consts/session-locals"
	"github.com/gofiber/fiber/v3"
)

func (h *MiddlewareHandler) PermissionsCheck(requiredPermissions []string) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {

		permissions, ok := c.Locals(sessionLocals.Permissions).([]string)
		if !ok {
			return c.
				Status(http.StatusForbidden).
				SendString(customerrorconst.FORBIDDEN_ERROR_MESSAGE)
		}

		if !arePermissionsSatisfied(requiredPermissions, permissions) {
			return c.
				Status(http.StatusForbidden).
				SendString(customerrorconst.FORBIDDEN_ERROR_MESSAGE)
		}

		return c.Next()
	}
}
