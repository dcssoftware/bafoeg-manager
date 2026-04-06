package customerrors

import (
	"fmt"
	"net/http"

	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

type VirusScannerError struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	error error
}

func NewVirusScannerError(
	err error,
) *VirusScannerError {

	file, lineNumber, _ := runtime.GetCallerFunction(1)
	errUuid := uuid.New()

	model := &VirusScannerError{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_DATABASE,
		httpUserMessage: errorconst.INTERNAL_SERVER_ERROR_MESSAGE,
		httpStatus:      http.StatusInternalServerError,

		error: err,
	}

	logger.ErrorWithCustomLocation(
		model.ID,
		model.httpUserMessage,
		model.file,
		model.lineNumber,
		model.error,
		"",
	)

	return model
}

func (e *VirusScannerError) Error() string {
	return e.error.Error()
}

func (e *VirusScannerError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *VirusScannerError) HTTPError() (int, string) {
	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
