package jwt

import (
	"testing"

	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
	"gotest.tools/assert"
)

func TestRefreshJWT(t *testing.T) {
	sessionUUID := "983ca250-b2ce-4fcb-bf3b-681d1e65ec7e"

	refreshToken, refreshTokenErr := CreateRefreshToken(provider.Authentik, sessionUUID)
	assert.NilError(t, refreshTokenErr, "could not create refresh token")

	refreshTokenData, refreshTokenDataErr := VerifyRefreshToken(refreshToken)
	assert.NilError(t, refreshTokenDataErr, "could not validate refresh token")

	assert.Equal(t, refreshTokenData.SessionID, sessionUUID, "session uuid does not match with jwt token claim")
}
