package customerrors

import (
	errorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
)

type SQLData map[string]any

type ErrorInterface interface {
	Error() (errormessage string)
	ErrorType() (errorIdentifier errorconst.ErrorIdentifier)
	HTTPError() (httpStatus int, userMessage string)
}

func GetHTTPError(err ErrorInterface) (int, string) {
	return err.HTTPError()
}
