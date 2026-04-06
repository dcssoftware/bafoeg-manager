package user

import (
	"testing"

	assets "github.com/dcssoftware/bafoeg-manager/integration-tests/resources/user/test-assets"
	integrationtestsetup "github.com/dcssoftware/bafoeg-manager/src/helper/integration-test-setup"
	"github.com/stretchr/testify/assert"
)

func TestUploadDefaultProfilePicture(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	user, userErr := insertBasicUser(t, testSetup)
	assert.NoError(t, userErr, "could not create user")

	testSetup.AppServices.UserSvc.UploadProfilePictureByID(nil, user.ID, []byte{})
}

func TestUploadEvilBenderProfilePicture(t *testing.T) {
	testSetup := integrationtestsetup.SetupTest(t)
	defer testSetup.Cleanup()

	user, userErr := insertBasicUser(t, testSetup)
	assert.NoError(t, userErr, "could not create user")

	oldPicture, oldPictureErr := testSetup.AppServices.UserSvc.GetProfilePictureByID(nil, user.ID)
	assert.NoError(t, oldPictureErr, "could not download the old profile picture")

	picture := assets.ProfilePictureEvilBender
	evilBenderErr := testSetup.AppServices.UserSvc.UploadProfilePictureByID(nil, user.ID, picture)
	assert.NoError(t, evilBenderErr, "could not upload evil bender profile picture")

	newPicture, newPictureErr := testSetup.AppServices.UserSvc.GetProfilePictureByID(nil, user.ID)
	assert.NoError(t, newPictureErr, "could not get evil bender profile picture after upload")
	assert.NotEqual(t, newPicture, []byte{}, "could not verify evil bender profile picture, picture is empty")
	assert.NotEqual(t, newPicture, oldPicture, "could not verify evil bender profile picture, picture is the same as old one")
}
