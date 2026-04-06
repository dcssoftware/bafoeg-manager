package customerrors

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"

	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/google/uuid"
)

type NotAuthorizedError struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	error error
}

func NewNotAuthorizedError() *NotAuthorizedError {

	var file string
	var lineNumber int

	_, callerFile, callerLine, ok := runtime.Caller(1)
	if ok {
		file = callerFile
		lineNumber = callerLine
	}

	errUuid := uuid.New()

	model := &NotAuthorizedError{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_UNAUTHORIZED,
		httpUserMessage: errorconst.NOT_AUTHORIZED_ERROR_MESSAGE,
		httpStatus:      http.StatusUnauthorized,

		error: errors.New("user is unauthorized"),
	}

	return model
}

func (e *NotAuthorizedError) Error() string {
	return e.error.Error()
}

func (e *NotAuthorizedError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *NotAuthorizedError) HTTPError() (int, string) {
	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
