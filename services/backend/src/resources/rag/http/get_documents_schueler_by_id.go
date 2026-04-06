package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/gofiber/fiber/v3"
)

func (h *RAGHandler) GetDocumentsSchülerByID(c fiber.Ctx) error {

	fileIDString := c.Params("fileID", "")
	if fileIDString == "" || !uuidvalidator.ValidateUUID(fileIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	document, documentModel, documentErr := h.service.GetDocumentSchülerByID(nil, fileIDString)
	if documentErr != nil {
		status, message := documentErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	c.Set("Content-Type", documentModel.FileType)
	return c.Status(http.StatusOK).Send(document)
}
