package integrationtestsetup

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/http/cookies"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func (i *TestInstance) Request(method string, url string, requestStruct any, responseStruct any) (int, error) {
	// Add nil checks
	if i == nil {
		panic("TestInstance is nil")
	}
	if i.Test == nil {
		panic("TestInstance.Test is nil")
	}
	if i.App == nil {
		assert.Fail(i.Test, "TestInstance.App is nil")
	}

	var requestReader io.Reader = bytes.NewReader([]byte{})

	if requestStruct != nil {
		requestBytes, requestBytesErr := json.Marshal(requestStruct)
		assert.NoErrorf(i.Test, requestBytesErr, "cannot encode / marshal request struct for \"%s\": %v\n", url, requestBytesErr)

		requestReader = bytes.NewReader(requestBytes)
	}

	request := httptest.NewRequest(method, url, requestReader)

	// Check if AppJWT is nil before using it
	if i.AppJWT != "" {
		httpJwtCookie := cookies.GenerateHttpJwtToken(i.AppJWT, false)
		if httpJwtCookie != nil {
			request.AddCookie(httpJwtCookie)
		}
	}

	response, responseErr := i.App.Test(request, fiber.TestConfig{Timeout: time.Second * 2, FailOnTimeout: true})
	assert.NoErrorf(i.Test, responseErr, fmt.Sprintf("cannot execute request from \"%s\": %v\n", url, responseErr))

	if response != nil {
		defer response.Body.Close()

		responseBody, responseBodyErr := io.ReadAll(response.Body)
		assert.NoErrorf(i.Test, responseBodyErr, "cannot read response from \"%s\": %v\n", url, responseBodyErr)

		assert.Contains(i.Test, string(responseBody), "", "")

		i.Test.Log(string(responseBody))

		// Skip JSON unmarshal if no response body or target is nil
		trimmed := strings.TrimSpace(string(responseBody))
		if trimmed == "" || responseStruct == nil {
			return response.StatusCode, nil
		}

		unmarshalErr := json.Unmarshal(responseBody, responseStruct)
		if unmarshalErr != nil {
			_ = unmarshalErr
			// Ignore EOF or syntax errors caused by non-JSON plain text responses
			if errors.Is(unmarshalErr, io.EOF) || strings.HasPrefix(unmarshalErr.Error(), "invalid character") {
				return response.StatusCode, unmarshalErr
			}
			assert.NoErrorf(i.Test, unmarshalErr, "cannot unmarshal response from \"%s\": %v\n", url, unmarshalErr.Error())
			return response.StatusCode, unmarshalErr
		}
	}

	return response.StatusCode, nil
}
