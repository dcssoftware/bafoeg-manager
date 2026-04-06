package middlewares

import (
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/gofiber/fiber/v3"
)

func (h *MiddlewareHandler) Logger() func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {

		endpointDurationStart := time.Now()

		nextHopErr := c.Next()

		endpointExecutionTime := time.Since(endpointDurationStart)

		logger.HTTP(
			c.Method(),
			c.OriginalURL(),
			c.Response().StatusCode(),
			endpointExecutionTime,
		)

		return nextHopErr
	}
}
