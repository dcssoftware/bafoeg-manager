package user

import (
	"testing"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	"github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
	"github.com/stretchr/testify/assert"
)

func insertBasicUser(t *testing.T, testSetup *integrationtestsetup.TestInstance) (*models.UserModel, customerrors.ErrorInterface) {
	userID, userErr := testSetup.AppServices.UserSvc.CreateUser(nil, "Test-User", "testusername", "test@example.com", true)
	assert.NoError(t, userErr, "could not create new user")

	return testSetup.AppServices.UserSvc.GetUserByID(nil, userID)
}
