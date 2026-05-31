package eakte

import (
	"testing"

	testassets "github.com/dcssoftware/bafoeg-manager/integration-tests/resources/eakte/test-assets"
	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	"github.com/stretchr/testify/assert"
)

func TestUploadEakte(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	uploadedEakteID, uploadedEakteIDErr := testSetup.AppServices.EakteSvc.UploadEakte(nil, testassets.XDomeaMessageStupid, "xdomea-message-stupid.zip")
	assert.NoError(t, uploadedEakteIDErr, "could not upload xdomea file")

	eakten, _, eaktenErr := testSetup.AppServices.EakteSvc.GetEakten(nil, 1)
	assert.NoError(t, eaktenErr, "could not retrieve eakten")
	assert.Len(t, eakten, 1, "should be exactly one eakte uploaded")

	documents, documentsCount, documentsErr := testSetup.AppServices.EakteSvc.GetFilesByAkteID(nil, uploadedEakteID.String())
	assert.NoError(t, documentsErr)
	assert.Equal(t, int(documentsCount), 3)
	assert.Len(t, documents, 3)

}
