package applications

import (
	"net/http"
	"testing"

	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	httpModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/http/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertFullApplications(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	_, newApplicant, appStatus := insertBasicApplication(t, testSetup)

	var applicationStructAfterInsert httpModel.ApplicationShortModels
	testSetup.Request(http.MethodGet, "/api/v1/applications", nil, &applicationStructAfterInsert)
	assert.NotEmpty(t, applicationStructAfterInsert.Application, "should be not empty: applications")

	assert.NotEqual(t, uuid.Nil.String(), applicationStructAfterInsert.Application[0].School.ID.String(), "not equal: school id")
	assert.NotEqual(t, uuid.Nil.String(), applicationStructAfterInsert.Application[0].AssignedUser.ID, "not equal: assigned user id")
	assert.Equal(t, newApplicant.ID, applicationStructAfterInsert.Application[0].Applicant.ID, "not equal: applicant id")
	assert.Equal(t, "10a 😅", applicationStructAfterInsert.Application[0].ClassLevel, "not equal: class level")
	assert.Equal(t, appStatus[0].Identifier, applicationStructAfterInsert.Application[0].Status.Identifier, "not equal: application status")
}
