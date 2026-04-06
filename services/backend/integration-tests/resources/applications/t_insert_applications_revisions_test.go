package applications

import (
	"testing"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	applicationsServiceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/stretchr/testify/assert"
)

func TestInsertApplicationRevision(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	applicationID, _, _ := insertBasicApplication(t, testSetup)
	_, applicationRevisionErr := testSetup.AppServices.ApplicationsService.InsertApplicationRevision(nil, applicationID.String(), "Latest Application Revision", "test test test", nil)
	assert.NoError(t, applicationRevisionErr, "could not insert new application revision")

	applicationRevisionsShort, applicationRevisionsShortErr := testSetup.AppServices.ApplicationsService.GetApplicationRevisionsByApplicationID(nil, 1, applicationID.String())
	assert.NoError(t, applicationRevisionsShortErr, "could not retrieve application revision short")
	assert.Equal(t, "Latest Application Revision", applicationRevisionsShort[0].Header, "header does not match")

	newDataModel := &applicationsServiceModel.ApplicationRevisionDataModel{}
	_, applicationRevisionErr2 := testSetup.AppServices.ApplicationsService.InsertApplicationRevision(nil, applicationID.String(), "Latest Application Revision", "test test test", newDataModel)
	assert.NoError(t, applicationRevisionErr2, "could not insert new application revision with new empty data")

	// todo insert new trainings address
	servideModel := &applicationsServiceModel.ApplicationRevisionTrainingsdata{
		Street:      "Musterstraße",
		HouseNumber: "123a",
		ZipCode:     "12345",
		City:        "Musterstadt",
		Country:     "Deutschland",
	}
	trainingsAddressID, trainingsAddressErr := testSetup.AppServices.ApplicationsStore.InsertTrainingsAddress(nil, servideModel)
	assert.NoError(t, trainingsAddressErr, "could not insert trainings address")

	newDataModel2 := &applicationsServiceModel.ApplicationRevisionDataModel{
		TrainingsAddressID: trainingsAddressID.String(),
	}
	_, applicationRevisionErr3 := testSetup.AppServices.ApplicationsService.InsertApplicationRevision(nil, applicationID.String(), "Latest Application Revision", "test test test", newDataModel2)
	assert.NoError(t, applicationRevisionErr3, "could not insert new application revision with new empty data")

	_ = applicationRevisionsShort
}
