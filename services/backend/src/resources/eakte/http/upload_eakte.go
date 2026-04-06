package http

import (
	"io"
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/gofiber/fiber/v3"
)

func (h *EakteHandler) UploadEakte(c fiber.Ctx) error {
	formFile, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE_MISSING_FILE)
	}

	file, fileErr := formFile.Open()
	if fileErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE_MISSING_FILE)
	}

	defer file.Close()

	fileContent, fileContentErr := io.ReadAll(file)
	if fileContentErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE_MISSING_FILE)
	}

	_, uploadErr := h.service.UploadEakte(nil, fileContent, formFile.Filename)
	if uploadErr != nil {
		httpStatus, userMessage := uploadErr.HTTPError()
		return c.Status(httpStatus).SendString(userMessage)
	}

	return c.Status(http.StatusOK).JSON("")
}
