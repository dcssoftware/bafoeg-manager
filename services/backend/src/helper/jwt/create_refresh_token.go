package jwt

import (
	"fmt"
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
	"github.com/golang-jwt/jwt/v5"
)

func CreateRefreshToken(authProvider provider.AuthProvider, sessionID string) (string, customerrors.ErrorInterface) {
	const hoursInDay = 24
	validityDays := time.Duration(configuration.Security.RefreshTokenValidityInDays) * hoursInDay

	claims := jwt.MapClaims{
		"auth-provider": authProvider,
		"session-uuid":  sessionID,
		"exp":           time.Now().Add(time.Hour * validityDays).Unix(),
		"authorized":    true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configuration.Security.RefreshTokenSigningKey))
	if err != nil {
		errMsg := fmt.Sprintf("Could not sign refresh token on creation %v", claims)
		return "", customerrors.NewInternalServerError(err, "", errMsg)
	}

	return tokenString, nil
}
