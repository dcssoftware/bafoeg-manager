package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	badrequestconstraints "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/bad-request-constraints"
	jwtHelper "github.com/dcssoftware/bafoeg-manager/src/helper/jwt"
	jwtHelperModels "github.com/dcssoftware/bafoeg-manager/src/helper/jwt/models"
	providerConst "github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
)

func (s *AuthService) CallbackFunction(provider, authCode, state string) (string, string, customerrors.ErrorInterface) {
	var loggedinUserID string
	var loggedinUserErr customerrors.ErrorInterface
	var providerFound bool = false

	for _, oauth := range s.oauthProvider {
		if provider == oauth.Identifier && oauth.Enabled {
			providerFound = true

			loggedinUserID, loggedinUserErr = oauth.AuthentikCallbackFunction(authCode, state)
			if loggedinUserErr != nil {
				return "", "", loggedinUserErr
			}

		}
	}

	if !providerFound {
		return "", "", customerrors.NewBadRequestError(badrequestconstraints.BadRequest_AuthenticationProviderInvalid)
	}

	tx, txErr := s.storage.StartTransaction()
	if txErr != nil {
		return "", "", txErr
	}

	defer func() {
		_ = tx.Rollback()
	}()

	// as long as they are not used, they are nil / null by default
	var dbIPAddr *string = nil
	var dbUserAgent *string = nil

	sessionID, sessionIDErr := s.storage.StartLoginSession(tx, loggedinUserID, dbUserAgent, dbIPAddr)
	if sessionIDErr != nil {
		return "", "", sessionIDErr
	}

	permissionScopes, permissionScopesErr := s.userSvc.GetUserPermissionsByID(tx, loggedinUserID)
	if permissionScopesErr != nil {
		return "", "", permissionScopesErr
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		return "", "", customerrors.NewDatabaseTransactionCommitError(commitErr, "Failed to commit transaction")
	}

	jwt, jwtErr := jwtHelper.CreateJWT(
		providerConst.Authentik,
		jwtHelperModels.NewJwtDataModel(loggedinUserID, sessionID, permissionScopes),
	)
	if jwtErr != nil {
		return "", "", customerrors.NewInternalServerError(jwtErr, loggedinUserID, "cannot create jwt after login / signup")
	}

	// refreshToken, refreshTokenErr := s.jwt.CreateRefreshToken(providerConst.Authentik, sessionID)
	refreshToken, refreshTokenErr := jwtHelper.CreateRefreshToken(providerConst.Authentik, sessionID)
	if refreshTokenErr != nil {
		return "", "", customerrors.NewInternalServerError(
			refreshTokenErr,
			loggedinUserID,
			"cannot create refreshToken after login / signup",
		)
	}

	return jwt, refreshToken, nil
}
