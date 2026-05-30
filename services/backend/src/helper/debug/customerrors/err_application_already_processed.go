package customerrors

import (
	"fmt"
	"net/http"

	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

type ApplicationAlreadyProcessedError struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int
}

func NewApplicationAlreadyProcessedError() *ApplicationAlreadyProcessedError {
	file, lineNumber, _ := runtime.GetCallerFunction(1)

	errUuid := uuid.New()

	model := &ApplicationAlreadyProcessedError{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_APPLICATION_ALREADY_PROCESSED,
		httpUserMessage: errorconst.APPLICATION_ALREADY_PROCESSED_ERROR_MESSAGE,
		httpStatus:      http.StatusBadRequest,
	}

	logger.ErrorWithCustomLocation(
		model.ID,
		model.httpUserMessage,
		model.file,
		model.lineNumber,
		nil,
		"",
	)

	return model
}

func (e *ApplicationAlreadyProcessedError) Error() string {
	_, message := e.HTTPError()
	return message
}

func (e *ApplicationAlreadyProcessedError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *ApplicationAlreadyProcessedError) HTTPError() (int, string) {
	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
