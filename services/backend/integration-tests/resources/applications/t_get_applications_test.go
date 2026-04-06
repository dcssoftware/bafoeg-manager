package applications

import (
	"net/http"
	"testing"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	"github.com/stretchr/testify/assert"
)

func TestGetApplications(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	var applicationStruct models.ApplicationShortModels
	testSetup.Request(http.MethodGet, "/api/v1/applications", nil, &applicationStruct)
	assert.Empty(t, applicationStruct.Application, "should be empty applications")
}
