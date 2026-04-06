package customerrors

import (
	"fmt"
	"net/http"

	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

type InternalServerError struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	errorContextData string

	requestingUserID string

	error error
}

func NewInternalServerError(err error, userID string, errorContextData string) *InternalServerError {

	file, lineNumber, _ := runtime.GetCallerFunction(1)
	errUuid := uuid.New()

	model := &InternalServerError{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_INTERNAL_SERVER_ERROR,
		httpUserMessage: errorconst.INTERNAL_SERVER_ERROR_MESSAGE,
		httpStatus:      http.StatusInternalServerError,

		errorContextData: errorContextData,

		requestingUserID: userID,

		error: err,
	}

	logger.ErrorWithCustomLocation(
		model.ID,
		model.httpUserMessage,
		model.file,
		model.lineNumber,
		model.error,
		errorContextData,
	)

	return model
}

func (e *InternalServerError) Error() string {
	return e.error.Error()
}

func (e *InternalServerError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *InternalServerError) HTTPError() (int, string) {
	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
