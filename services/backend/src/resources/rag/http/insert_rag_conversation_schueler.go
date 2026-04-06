package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/http/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/rag/service/consts"
	sessionLocals "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/http/consts/session-locals"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *RAGHandler) InsertRagConversationSchueler(c fiber.Ctx) error {
	userIDString := c.Locals(sessionLocals.UserUUID).(string)
	userID, userIDErr := uuid.Parse(userIDString)
	if userIDErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	conversationID, conversationIDErr := h.service.InsertRagConversation(nil, userID, consts.BafögTypeSchueler)
	if conversationIDErr != nil {
		return c.Status(http.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
	}

	return c.JSON(models.ToHttpRAGConversationIDResponseModel(conversationID))
}
