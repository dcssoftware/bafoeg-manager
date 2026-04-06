package applications

import (
	"net/http"
	"testing"

	testassets "github.com/dcssoftware/bafoeg-manager/integration-tests/resources/applications/test-assets"
	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	"github.com/stretchr/testify/assert"
)

func TestUploadApplicationFiles(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	applicationID, _, _ := insertBasicApplication(t, testSetup)

	insert1Err := testSetup.AppServices.ApplicationsService.UploadApplicationFile(nil, applicationID.String(), "example01.pdf", testassets.ExampleDocumentPDF01)
	assert.NoError(t, insert1Err)

	var applicationStruct *models.ApplicationFileModels
	testSetup.Request(http.MethodGet, "/api/v1/applications/"+applicationID.String()+"/files", nil, &applicationStruct)
	assert.Equal(t, uint(1), applicationStruct.MaxCount, "should be 1 application file maxcount")
	assert.Len(t, applicationStruct.Files, 1, "should be 1 application file")
	assert.Equal(t, float64(35339), applicationStruct.Files[0].File.Size, "should be file size 35339")

	insert2Err := testSetup.AppServices.ApplicationsService.UploadApplicationFile(nil, applicationID.String(), "example02.pdf", testassets.ExampleDocumentPDF02)
	assert.NoError(t, insert2Err)

	var applicationStruct2 *models.ApplicationFileModels
	testSetup.Request(http.MethodGet, "/api/v1/applications/"+applicationID.String()+"/files", nil, &applicationStruct2)
	assert.Equal(t, uint(2), applicationStruct2.MaxCount, "should be 2 application file maxcount")
	assert.Len(t, applicationStruct2.Files, 2, "should be 2 application files")
	assert.Equal(t, float64(966501), applicationStruct2.Files[1].File.Size, "should be file size 966501")

}
