package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *EakteHandler) GetDocumentFileByID(c fiber.Ctx) error {

	documentIDString := c.Params("documentID", "")
	if documentIDString == "" || !uuidvalidator.ValidateUUID(documentIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	_, file, err := h.service.GetDocumentFileByID(nil, uuid.MustParse(documentIDString))
	if err != nil {
		status, message := err.HTTPError()
		return c.Status(status).SendString(message)
	}

	c.Set("Content-Type", "application/pdf")
	return c.Status(http.StatusOK).Send(file)
}
