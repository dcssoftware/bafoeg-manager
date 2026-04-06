package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	providerConst "github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
)

/*
IMPORTANT!
A login is a connection user <-> auth provider (e.g. authentik, keycloak, etc.) not an active session.
The session is stored separately.
*/
func (s *UserService) GetUserIDByLogin(
	providerType providerConst.AuthProvider,
	providerID string,
) (
	string,
	customerrors.ErrorInterface,
) {

	return s.storage.GetIDByLogin(string(providerType), providerID)

}
