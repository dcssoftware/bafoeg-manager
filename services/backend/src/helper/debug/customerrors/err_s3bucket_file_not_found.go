package customerrors

import (
	"fmt"
	"net/http"

	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/runtime"
	"github.com/google/uuid"
)

type S3BucketDownloadFileNotFoundError struct {
	ID *uuid.UUID

	file       string
	lineNumber int

	errorIdentifier errorconst.ErrorIdentifier
	httpUserMessage string
	httpStatus      int

	errorContextData string
	error            error
}

func NewS3BucketDownloadFileNotFoundError(err error, errorContextData string) *S3BucketDownloadFileNotFoundError {
	file, lineNumber, _ := runtime.GetCallerFunction(1)
	errUuid := uuid.New()

	model := &S3BucketDownloadFileNotFoundError{
		ID: &errUuid,

		file:       file,
		lineNumber: lineNumber,

		errorIdentifier: errorconst.ERROR_IDENTIFIER_S3_NOT_FOUND,
		httpUserMessage: errorconst.S3BUCKET_NOT_FOUND_ERROR_MESSAGE,
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

func (e *S3BucketDownloadFileNotFoundError) Error() string {
	return e.error.Error()
}

func (e *S3BucketDownloadFileNotFoundError) ErrorType() errorconst.ErrorIdentifier {
	return e.errorIdentifier
}

func (e *S3BucketDownloadFileNotFoundError) HTTPError() (int, string) {
	return e.httpStatus, fmt.Sprintf("[ %s ]: (#%s)<br>%s", e.errorIdentifier, e.ID, e.httpUserMessage)
}
