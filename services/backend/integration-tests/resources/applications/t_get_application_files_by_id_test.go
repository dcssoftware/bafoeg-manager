package applications

import (
	"net/http"
	"testing"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	"github.com/stretchr/testify/assert"
)

func TestGetApplicationFilesByApplicationID(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	applicationID, _, _ := insertBasicApplication(t, testSetup)

	var applicationStruct *models.ApplicationFileModels
	testSetup.Request(http.MethodGet, "/api/v1/applications/"+applicationID.String()+"/files", nil, &applicationStruct)
	assert.Empty(t, applicationStruct.Files, "should be empty applications")
}
