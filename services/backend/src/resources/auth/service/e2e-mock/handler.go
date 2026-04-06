package e2emock

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"

	providerConst "github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
	userService "github.com/dcssoftware/bafoeg-manager/src/resources/user/service"
)

type AuthService struct {
	Enabled    bool
	Identifier string

	userSvc *userService.UserService
}

func NewAuthService(userSvc *userService.UserService) *AuthService {
	return &AuthService{
		Enabled:    true,
		Identifier: providerConst.E2EMock.String(),

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
