package jwt

import (
	"fmt"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	models "github.com/dcssoftware/bafoeg-manager/src/helper/jwt/models"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(jwtString string) (*models.JwtDataModel, error) {

	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (any, error) {
		_, okECDSA := token.Method.(*jwt.SigningMethodECDSA)
		_, okHS256 := token.Method.(*jwt.SigningMethodHMAC)

		if !okECDSA && !okHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configuration.Security.JwtTokenSigningKey), nil
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

	var permissionScopes []string
	if value, ok := claims["scopes"].([]any); ok {
		for _, v := range value {
			if str, ok := v.(string); ok {
				permissionScopes = append(permissionScopes, str)
			}
		}
	}

	userUUID, userUUIDok := claims["user-uuid"].(string)
	if !userUUIDok {
		return nil, fmt.Errorf("unauthorized")
	}

	sessionUUID, sessionUUIDok := claims["session-uuid"].(string)
	if !sessionUUIDok {
		return nil, fmt.Errorf("unauthorized")
	}

	data := models.NewJwtDataModel(
		userUUID,
		sessionUUID,
		permissionScopes,
	)

	return data, nil
}
