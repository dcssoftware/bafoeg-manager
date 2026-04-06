package http

import (
	"io"
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	sessionlocals "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/http/consts/session-locals"
	"github.com/gofiber/fiber/v3"
	"github.com/h2non/filetype"
)

func (h *RAGHandler) UploadRAGrelevantDocumentsStudierende(c fiber.Ctx) error {
	pdfFile, pdfFileErr := c.FormFile("file")
	if pdfFileErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE_MISSING_FILE)
	}

	userID, ok := c.Locals(sessionlocals.UserUUID).(string)
	if !ok {
		return c.Status(http.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
	}

	multipartFile, multipartFileErr := pdfFile.Open()
	if multipartFileErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE_MISSING_FILE)
	}

	defer multipartFile.Close()

	fileContent, fileContentErr := io.ReadAll(multipartFile)
	if fileContentErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	kind, kindErr := filetype.Match(fileContent)
	if kindErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	} else if kind.MIME.Value != "application/pdf" {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE_NOT_PDF_FILE)
	}

	err := h.service.UploadRAGrelevantDocumentsPDStudierende(nil, pdfFile.Filename, kind.MIME.Value, pdfFile.Size, fileContent, userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
	}

	return c.SendStatus(http.StatusCreated)
}
