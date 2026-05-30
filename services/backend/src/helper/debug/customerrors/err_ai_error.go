package customerrors

import (
	"fmt"
	"net/http"

	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

type AIError struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	prompt string
	error  error
}

func NewAIError(err error, prompt string) *AIError {
	file, lineNumber, _ := runtime.GetCallerFunction(1)

	errUuid := uuid.New()

	model := &AIError{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_AI_NO_RESPONSE,
		httpUserMessage: errorconst.AI_NO_RESPONSE_ERROR_MESSAGE,
		httpStatus:      http.StatusInternalServerError,

		prompt: prompt,
		error:  err,
	}

	logger.ErrorWithCustomLocation(
		model.ID,
		model.httpUserMessage,
		model.file,
		model.lineNumber,
		model.error,
		prompt,
	)

	return model
}

func (e *AIError) Error() string {
	return e.error.Error()
}

func (e *AIError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *AIError) HTTPError() (int, string) {
	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
