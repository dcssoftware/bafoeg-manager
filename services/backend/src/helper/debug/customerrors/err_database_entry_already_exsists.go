package customerrors

import (
	"fmt"
	"net/http"

	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

type DatabaseEntryAlreadyExistsErr struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	error error
}

func NewDatabaseEntryAlreadyExistsErr() *DatabaseEntryAlreadyExistsErr {
	file, lineNumber, _ := runtime.GetCallerFunction(1)

	errUuid := uuid.New()

	model := &DatabaseEntryAlreadyExistsErr{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_DATABASE_CONFLICT,
		httpUserMessage: errorconst.DATABASE_CONFLICT_ERROR_MESSAGE,
		httpStatus:      http.StatusConflict,

		error: fmt.Errorf(
			"database conflict [key: %s] %s",
			errorconst.ERROR_IDENTIFIER_DATABASE_CONFLICT,
			errorconst.DATABASE_CONFLICT_ERROR_MESSAGE,
		),
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

func (e *DatabaseEntryAlreadyExistsErr) Error() string {
	return e.error.Error()
}

func (e *DatabaseEntryAlreadyExistsErr) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *DatabaseEntryAlreadyExistsErr) HTTPError() (int, string) {
	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
