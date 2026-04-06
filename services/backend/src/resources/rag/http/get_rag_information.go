package http

import (
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/http/models"
	"github.com/gofiber/fiber/v3"
)

func (h *RAGHandler) GetRagInformation(c fiber.Ctx) error {

	var response models.RagInformationModel = models.RagInformationModel{
		AIModelServerAddress: configuration.OllamaAPI.Address,
		AIModelServerPort:    configuration.OllamaAPI.Port,
		AIModelServerSecure:  configuration.OllamaAPI.IsSecure,

		EmbeddingModelname:  configuration.OllamaAPI.EmbeddingModelname,
		RequestingModelname: configuration.OllamaAPI.RequestingModelname,
	}

	return c.JSON(response)
}
