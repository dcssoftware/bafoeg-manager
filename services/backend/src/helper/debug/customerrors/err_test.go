package customerrors

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/google/uuid"
	"gotest.tools/assert"
)

func TestHttpDatabaseErrorHandling(t *testing.T) {
	logger.NewLogger()

	myDatabaseError := NewDatabaseError(
		errors.New("Database error"),
		uuid.New().String(),
		"SELECT * FROM users",
		"",
		map[string]any{"id": 123},
	)

	httpStatus, httpUserMessage := GetHTTPError(myDatabaseError)

	assert.Equal(t, httpStatus, http.StatusInternalServerError)
	assert.Assert(
		t,
		strings.Contains(httpUserMessage, customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE),
		"Errormessage does not contain defined error type",
	)
}

func TestHttpInternalServerErrorHandling(t *testing.T) {
	logger.NewLogger()

	myDatabaseError := NewInternalServerError(
		errors.New("Internal server error"),
		uuid.New().String(),
		"Error context data",
	)

	httpStatus, httpUserMessage := GetHTTPError(myDatabaseError)

	assert.Equal(t, httpStatus, http.StatusInternalServerError)

	assert.Assert(
		t,
		strings.Contains(httpUserMessage, customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE),
		"Errormessage does not contain defined error type",
	)
}
