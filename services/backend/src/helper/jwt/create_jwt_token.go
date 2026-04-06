package jwt

import (
	"fmt"
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	models "github.com/dcssoftware/bafoeg-manager/src/helper/jwt/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
	"github.com/golang-jwt/jwt/v5"
)

// split for testing purposes
func CreateJWT(authProvider provider.AuthProvider, jwtData *models.JwtDataModel) (string, customerrors.ErrorInterface) {
	validityMinutes := time.Duration(configuration.Security.JwtTokenValidInMinutes)
	expires := time.Now().Add(validityMinutes * time.Minute).Unix()
	return createJWT(authProvider, jwtData, expires)
}

func createJWT(
	authProvider provider.AuthProvider,
	jwtData *models.JwtDataModel,
	expires int64,
) (
	string,
	customerrors.ErrorInterface,
) {
	claims := jwt.MapClaims{
		"auth-provider": authProvider,
		"user-uuid":     jwtData.UserUUID,
		"session-uuid":  jwtData.SessionID,
		"scopes":        jwtData.Scopes,
		"exp":           expires,
		"authorized":    true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configuration.Security.JwtTokenSigningKey))
	if err != nil {
		errMsg := fmt.Sprintf("Could not sign jwt on creation %v", claims)
		return "", customerrors.NewInternalServerError(err, jwtData.UserUUID, errMsg)
	}

	return tokenString, nil
}
