package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/http/models"
	"github.com/gofiber/fiber/v3"
)

func (h *RAGHandler) GetDocumentsStudierenden(c fiber.Ctx) error {

	queries := c.Queries()
	filterResultString := queries["filterResult"]
	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(queries)
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	documents, documentsErr := h.service.GetDocumentsStudierenden(nil, pageNumber, filterResultString)
	if documentsErr != nil {
		status, message := documentsErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	documentsCount, documentsCountErr := h.service.GetDocumentsStudierendenCount(nil, filterResultString)
	if documentsCountErr != nil {
		status, message := documentsCountErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	responseModel := models.ToDocumentsHttpModels(documents, documentsCount)

	return c.Status(http.StatusOK).JSON(responseModel)
}
