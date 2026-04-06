package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	authentikProviderService "github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/authentik"
	userService "github.com/dcssoftware/bafoeg-manager/src/resources/user/service"
	"github.com/go-sqlx/sqlx"
)

type AuthService struct {
	storage AuthStorageInterface

	userSvc *userService.UserService

	oauthProvider []*authentikProviderService.AuthService
}

func NewAuthService(storage AuthStorageInterface, userService *userService.UserService) *AuthService {
	var oauthServices []*authentikProviderService.AuthService

	for _, oauth := range configuration.Authentication.Oauth {
		oauthService := authentikProviderService.NewAuthService(userService, oauth)
		oauthServices = append(oauthServices, oauthService)

	}

	return &AuthService{
		storage: storage,
		userSvc: userService,

		oauthProvider: oauthServices,
	}
}

type AuthStorageInterface interface {
	StartTransaction() (tx *sqlx.Tx, err customerrors.ErrorInterface)
	StartLoginSession(
		tx *sqlx.Tx,
		loggedinUser string,
		useragentHash,
		ipaddr *string) (sessionID string, err customerrors.ErrorInterface)
}
