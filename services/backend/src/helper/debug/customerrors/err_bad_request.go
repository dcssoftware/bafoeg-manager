package customerrors

import (
	"fmt"
	"net/http"

	badrequestconstraints "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/bad-request-constraints"
	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

type BadRequestError struct {
	ID string

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	inputKey   string
	constraint string

	error error
}

func NewBadRequestError(badRequestContraint *badrequestconstraints.BadRequestContraint) *BadRequestError {

	file, lineNumber, _ := runtime.GetCallerFunction(1)

	model := &BadRequestError{
		ID: uuid.New().String(),

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_BAD_REQUEST,
		httpUserMessage: errorconst.BAD_REQUEST_ERROR_MESSAGE,
		httpStatus:      http.StatusBadRequest,

		inputKey:   badRequestContraint.KeyName,
		constraint: badRequestContraint.ConstraintMessage,

		error: fmt.Errorf("bad request [key: %s] %s", badRequestContraint.KeyName, badRequestContraint.ConstraintMessage),
	}

	// logger.InfoWithCustomLocation(
	// 	model.ID,
	// 	model.httpUserMessage,
	// 	model.file,
	// 	model.lineNumber,
	// 	nil,
	// )

	return model
}

func (e *BadRequestError) Error() string {
	return e.error.Error()
}

func (e *BadRequestError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *BadRequestError) HTTPError() (int, string) {

	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
