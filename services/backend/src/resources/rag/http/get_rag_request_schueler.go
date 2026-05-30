package http

import (
	"bufio"
	"context"
	"net/http"
	"strings"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	sessionLocals "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/http/consts/session-locals"
	"github.com/gofiber/fiber/v3"
	"github.com/goforj/godump"
	"github.com/google/uuid"
)

func (h *RAGHandler) GetRAGrequestSchüler(c fiber.Ctx) error {
	userIDString := c.Locals(sessionLocals.UserUUID).(string)
	userID, userIDErr := uuid.Parse(userIDString)
	if userIDErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	prompt := c.Query("prompt", "")
	if strings.TrimSpace(prompt) == "" {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE_MISSING_PROMPT)
	}

	conversationIDString := c.Query("conversationID", uuid.Nil.String())
	conversationID, conversationIDErr := uuid.Parse(conversationIDString)
	if conversationIDErr != nil {
		conversationID = uuid.Nil
	}

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	c.RequestCtx().SetBodyStreamWriter(func(w *bufio.Writer) {

		_, _, responseErr := h.service.GetRAGrequestSchüler(
			nil,
			conversationID,
			userID,
			prompt,
			func(ctx context.Context, reasoningChunk []byte, chunk []byte) error {

				godump.Dump("HERE")

				if _, writeErr := w.Write(reasoningChunk); writeErr != nil {
					return writeErr
				}

				if _, writeErr := w.Write([]byte("\n----------\n")); writeErr != nil {
					return writeErr
				}

				// Write each chunk to the response stream
				if _, writeErr := w.Write(chunk); writeErr != nil {
					return writeErr
				}
				// Flush to ensure data is sent immediately
				return w.Flush()
			},
		)
		if responseErr != nil {
			_, message := responseErr.HTTPError()
			w.Write([]byte("event: error\ndata: " + message + "\n\n"))
			w.Flush()
		}

	})
	return nil
}
