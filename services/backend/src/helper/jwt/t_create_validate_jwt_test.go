package jwt

import (
	"testing"
	"time"

	jwt_models "github.com/dcssoftware/bafoeg-manager/src/helper/jwt/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndValidateJWT(t *testing.T) {
	userUUID := "1e3896c6-b4c0-4e82-95bf-76ed6b369c40"
	sessionUUID := "983ca250-b2ce-4fcb-bf3b-681d1e65ec7e"
	scopes := []string{"test", "jwt", "now!"}

	jwtModelData := jwt_models.NewJwtDataModel(userUUID, sessionUUID, scopes)
	jwtToken, jwtTokenErr := CreateJWT(provider.Authentik, jwtModelData)
	assert.NoError(t, jwtTokenErr, "could not create jwt token")

	jwtData, jwtDataErr := VerifyJWT(jwtToken)
	assert.NoError(t, jwtDataErr, "could not validate jwt token")

	assert.Equal(t, jwtData.UserUUID, userUUID, "user uuid does not match with jwt token claim")
	assert.Equal(t, jwtData.SessionID, sessionUUID, "session uuid does not match with jwt token claim")

	// todo
	// assert.Equal(t, len(jwtData.Scopes), len(scopes), "scopes array size does not match with jwt token claim")
	// assert.DeepEqual(t, jwtData.Scopes, scopes)
}

func TestValidateExpiredJWT(t *testing.T) {
	userUUID := "1e3896c6-b4c0-4e82-95bf-76ed6b369c40"
	sessionUUID := "983ca250-b2ce-4fcb-bf3b-681d1e65ec7e"
	scopes := []string{"test", "jwt", "now!"}

	jwtModelData := jwt_models.NewJwtDataModel(userUUID, sessionUUID, scopes)
	jwtToken, jwtTokenErr := createJWT(provider.Authentik, jwtModelData, time.Now().Add(time.Minute*-5).Unix())
	assert.NoError(t, jwtTokenErr, "could not create jwt token")

	_, jwtDataErr := VerifyJWT(jwtToken)
	assert.ErrorContains(t, jwtDataErr, "expired", "could not validate that invalid jwt was successfully flagged as expired")
}
