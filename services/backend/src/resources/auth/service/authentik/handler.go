package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"golang.org/x/oauth2"

	userService "github.com/dcssoftware/bafoeg-manager/src/resources/user/service"
)

type AuthService struct {
	authentikConfig          *oauth2.Config
	authentikOauthUserInfoEP string
	authentikOauthLogoutEP   string

	Enabled    bool
	Identifier string

	userSvc *userService.UserService
}

func NewAuthService(userSvc *userService.UserService, oauth configuration.OauthConfigurationModel) *AuthService {
	authentikConfig := &oauth2.Config{
		ClientID:     oauth.AccessKeyID,
		ClientSecret: oauth.AccessSecretID,

		RedirectURL: oauth.CallbackURL,
		Scopes:      oauth.Scopes,

		Endpoint: oauth2.Endpoint{
			AuthURL:  oauth.AuthorizationURL,
			TokenURL: oauth.TokenURL,

			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	return &AuthService{
		Enabled:    oauth.Enabled,
		Identifier: oauth.Identifier,

		authentikConfig:          authentikConfig,
		authentikOauthUserInfoEP: oauth.UserinfoURL,
		authentikOauthLogoutEP:   oauth.LogoutURL,

		userSvc: userSvc,
	}
}

type AuthStorageInterface interface {
	StartTransaction() (*sqlx.Tx, customerrors.ErrorInterface)
	GetUserIDByLogin(
		provider string,
		providerID string) (string, customerrors.ErrorInterface)
	CreateUser(
		tx *sqlx.Tx,
		displayName string,
		username string,
		email string,
		emailVerified bool) (string, customerrors.ErrorInterface)
	CreateUserLoginIdentity(
		tx *sqlx.Tx,
		createdUserID string,
		authProvider string,
		authProviderID string) customerrors.ErrorInterface
}
