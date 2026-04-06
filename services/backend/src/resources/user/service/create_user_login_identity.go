package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	providerConst "github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
	"github.com/go-sqlx/sqlx"
)

func (s *UserService) CreateUserLoginIdentity(
	tx *sqlx.Tx,
	createdUserID string,
	providerType providerConst.AuthProvider,
	providerID string,
) customerrors.ErrorInterface {
	return s.storage.CreateLoginIdentity(tx, createdUserID, string(providerType), providerID)
}
