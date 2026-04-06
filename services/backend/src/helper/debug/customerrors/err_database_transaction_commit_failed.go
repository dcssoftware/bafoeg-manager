package customerrors

import (
	"fmt"
	"net/http"

	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

type DatabaseTransactionCommitError struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	errorContextData string

	error error
}

type DatabaseTransactionRollbackError struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	errorContextData string

	error error
}

func NewDatabaseTransactionCommitError(err error, errorContextData string) *DatabaseTransactionCommitError {

	file, lineNumber, _ := runtime.GetCallerFunction(1)
	errUuid := uuid.New()

	model := &DatabaseTransactionCommitError{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_DATABASE_TRANSACTION_NO_COMMIT,
		httpUserMessage: errorconst.INTERNAL_SERVER_ERROR_MESSAGE,
		httpStatus:      http.StatusInternalServerError,

		errorContextData: errorContextData,

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

func NewDatabaseTransactionRollbackError(err error, errorContextData string) *DatabaseTransactionRollbackError {

	file, lineNumber, _ := runtime.GetCallerFunction(1)
	errUuid := uuid.New()

	model := &DatabaseTransactionRollbackError{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_DATABASE_TRANSACTION_NO_ROLLBACK,
		httpUserMessage: errorconst.INTERNAL_SERVER_ERROR_MESSAGE,
		httpStatus:      http.StatusInternalServerError,

		errorContextData: errorContextData,

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

func (e *DatabaseTransactionCommitError) Error() string {
	return e.error.Error()
}

func (e *DatabaseTransactionCommitError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *DatabaseTransactionCommitError) HTTPError() (int, string) {

	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
