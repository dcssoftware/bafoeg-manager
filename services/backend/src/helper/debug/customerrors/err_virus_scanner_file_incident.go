package customerrors

import (
	"fmt"
	"net/http"

	badrequestconstraints "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/bad-request-constraints"
	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

type VirusScannerIncidentError struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int
	error           error
}

func NewVirusScannerIncidentError() *VirusScannerIncidentError {

	file, lineNumber, _ := runtime.GetCallerFunction(1)
	errUuid := uuid.New()

	model := &VirusScannerIncidentError{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_BAD_REQUEST,
		httpUserMessage: errorconst.VIRUS_SCAN_BLOCKED_ERROR_MESSAGE,
		httpStatus:      http.StatusBadRequest,

		error: fmt.Errorf(
			"bad request [key: %s] %s",
			badrequestconstraints.BadRequest_FileInfected.KeyName,
			badrequestconstraints.BadRequest_FileInfected.ConstraintMessage,
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

func (e *VirusScannerIncidentError) Error() string {
	return e.error.Error()
}

func (e *VirusScannerIncidentError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *VirusScannerIncidentError) HTTPError() (int, string) {
	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
