package applications

import (
	"net/http"
	"testing"

	assets "github.com/dcssoftware/bafoeg-manager/integration-tests/resources/applications/test-assets"
	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	"github.com/stretchr/testify/assert"
)

func TestGetApplicationFilesWithEakteFiles(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	applicationID, _, _ := insertBasicApplication(t, testSetup)

	var applicationStruct *models.ApplicationFileModels
	testSetup.Request(http.MethodGet, "/api/v1/applications/"+applicationID.String()+"/files", nil, &applicationStruct)
	assert.Empty(t, applicationStruct.Files, "should be empty applications")

	akteID, uploadAkteErr := testSetup.AppServices.EakteSvc.UploadEakte(nil, assets.XDomeaMessageStupid, "xdomea-message-stupid.zip")
	assert.NoError(t, uploadAkteErr)

	testSetup.AppServices.ApplicationsService.InsertApplicationEakteMapping(nil, applicationID, akteID)

	fileCount, fileCountErr := testSetup.AppServices.ApplicationsService.GetApplicationFilesByApplicationIDCount(nil, applicationID.String())
	assert.NoError(t, fileCountErr)
	assert.Equal(t, 3, int(fileCount))

	files, fileErr := testSetup.AppServices.ApplicationsService.GetApplicationFilesByApplicationID(nil, 1, applicationID.String())
	assert.NoError(t, fileErr)
	assert.Len(t, files, 3)
}
