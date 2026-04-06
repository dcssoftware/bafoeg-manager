package customerrors

import (
	"fmt"
	"net/http"

	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

type S3BucketUploadFailedError struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	errorContextData string
	error            error
}

func NewS3BucketUploadFailedError(err error, errorContextData string) *S3BucketUploadFailedError {
	file, lineNumber, _ := runtime.GetCallerFunction(1)
	errUuid := uuid.New()

	model := &S3BucketUploadFailedError{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_S3_NOT_UPLOADED,
		httpUserMessage: errorconst.S3BUCKET_NOT_UPLOADED_ERROR_MESSAGE,
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

func (e *S3BucketUploadFailedError) Error() string {
	return e.error.Error()
}

func (e *S3BucketUploadFailedError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *S3BucketUploadFailedError) HTTPError() (int, string) {
	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
