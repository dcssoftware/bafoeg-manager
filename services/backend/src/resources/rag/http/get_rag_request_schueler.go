package http

import (
	"bufio"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	sessionLocals "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/http/consts/session-locals"
	"github.com/gofiber/fiber/v3"
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

	requestCtx := c.RequestCtx()
	ctx, cancel := context.WithTimeout(requestCtx, time.Second*time.Duration(configuration.OllamaAPI.RequestTimeoutSeconds))

	requestCtx.SetBodyStreamWriter(func(w *bufio.Writer) {
		defer cancel()

		_, _, responseErr := h.service.GetRAGrequestSchüler(
			ctx,
			nil,
			conversationID,
			userID,
			prompt,
			func(ctx context.Context, chunk []byte) error {

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

// fiber:context-methods migrated
