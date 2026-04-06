package applications

import (
	"testing"

	testassets "github.com/dcssoftware/bafoeg-manager/integration-tests/resources/applications/test-assets"
	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	"github.com/stretchr/testify/assert"
)

func TestUploadEicarInfectedApplicationFiles(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	applicationID, _, _ := insertBasicApplication(t, testSetup)

	infectedTxtErr := testSetup.AppServices.ApplicationsService.UploadApplicationFile(nil, applicationID.String(), "", testassets.ExampleDocumentTxtInfected)
	assert.Contains(t, infectedTxtErr.Error(), "The given file contains a known virus", "cannot detect virus on file upload (txt)")

	infectedZip1Err := testSetup.AppServices.ApplicationsService.UploadApplicationFile(nil, applicationID.String(), "", testassets.ExampleDocumentZipInfected01)
	assert.Contains(t, infectedZip1Err.Error(), "The given file contains a known virus", "cannot detect virus on file upload (zip)")

	infectedZip2Err := testSetup.AppServices.ApplicationsService.UploadApplicationFile(nil, applicationID.String(), "", testassets.ExampleDocumentTxtInfected02)
	assert.Contains(t, infectedZip2Err.Error(), "The given file contains a known virus", "cannot detect virus on file upload (zip)")
}
