package jwt

import (
	"fmt"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	models "github.com/dcssoftware/bafoeg-manager/src/helper/jwt/models"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyRefreshToken(refreshTokenString string) (*models.RefreshDataModel, error) {

	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (any, error) {
		_, okECDSA := token.Method.(*jwt.SigningMethodECDSA)
		_, okHS256 := token.Method.(*jwt.SigningMethodHMAC)

		if !okECDSA && !okHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configuration.Security.RefreshTokenSigningKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims["authorized"] != true {
		return nil, fmt.Errorf("unauthorized")
	}

	var sessionUUID string
	if val, ok := claims["session-uuid"]; ok {
		sessionUUID, ok = val.(string)
		if !ok {
			return nil, fmt.Errorf("could not read session-uuid")
		}
	}

	data := models.NewRefreshDataModel(sessionUUID)
	return data, nil
}
