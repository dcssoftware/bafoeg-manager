package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	custombadrequestconstraints "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/bad-request-constraints"
)

func (s *AuthService) CreateRedirect(providerIdentifier string, state string) (string, customerrors.ErrorInterface) {
	for _, oauth := range s.oauthProvider {
		if providerIdentifier == oauth.Identifier && oauth.Enabled {
			return oauth.CreateRedirect(state), nil
		}
	}
	return "", customerrors.NewBadRequestError(custombadrequestconstraints.BadRequest_AuthenticationProviderInvalid)
}
