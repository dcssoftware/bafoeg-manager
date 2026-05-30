package customerrors

import (
	"fmt"
	"net/http"

	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

type DatabaseNotFoundError struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	requestingUserID string

	sqlQuery string
	sqlData  map[string]any

	error error
}

func NewDatabaseNotFoundError(err error, userID string, sqlQuery string, sqlData SQLData) *DatabaseNotFoundError {
	file, lineNumber, _ := runtime.GetCallerFunction(1)

	errUuid := uuid.New()

	model := &DatabaseNotFoundError{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_DATABASE_NOT_FOUND,
		httpUserMessage: errorconst.DATABASE_NOT_FOUND_ERROR_MESSAGE,
		httpStatus:      http.StatusNotFound,

		requestingUserID: userID,

		sqlQuery: sqlQuery,
		sqlData:  sqlData,

		error: err,
	}

	logger.ErrorWithCustomLocation(
		model.ID,
		model.httpUserMessage,
		model.file,
		model.lineNumber,
		model.error,
		fmt.Sprintf("%v ;; %v", model.sqlQuery, model.sqlData),
	)

	return model
}

func (e *DatabaseNotFoundError) Error() string {
	return e.error.Error()
}

func (e *DatabaseNotFoundError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *DatabaseNotFoundError) HTTPError() (int, string) {
	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
