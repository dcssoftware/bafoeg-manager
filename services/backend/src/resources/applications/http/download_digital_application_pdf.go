package http

import "github.com/gofiber/fiber/v3"

func (h *ApplicationsHandler) DownloadDigitalApplicationPDF(c fiber.Ctx) error {

	documents, documentsErr := h.service.DownloadDigitalApplicationPDF()
	if documentsErr != nil {
		status, message := documentsErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.Send(documents)
}
