package customerrors

import (
	"fmt"
	"net/http"

	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

type S3BucketNoRemoveError struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	errorContextData string
	error            error
}

func NewS3BucketNoRemoveError(err error, errorContextData string) *S3BucketNoRemoveError {
	file, lineNumber, _ := runtime.GetCallerFunction(1)
	errUuid := uuid.New()

	model := &S3BucketNoRemoveError{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_S3_NOT_DELETED,
		httpUserMessage: errorconst.S3BUCKET_NOT_DELETED_ERROR_MESSAGE,
		httpStatus:      http.StatusInternalServerError,

		errorContextData: errorContextData,

		error: err,
	}

	return model
}

func (e *S3BucketNoRemoveError) Error() string {
	return e.error.Error()
}

func (e *S3BucketNoRemoveError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *S3BucketNoRemoveError) HTTPError() (int, string) {
	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
