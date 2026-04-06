package http

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func (h *RAGHandler) GetRAGrequestStudierende(c fiber.Ctx) error {
	prompt := c.Query("prompt", "")

	if strings.TrimSpace(prompt) == "" {
		return c.Status(http.StatusBadRequest).SendString("")
	}

	response, responseErr := h.service.GetRAGrequestStudierenden(prompt)
	if responseErr != nil {
		return c.Status(http.StatusInternalServerError).SendString(responseErr.Error())
	}

	return c.Status(http.StatusOK).SendString(response)
}
