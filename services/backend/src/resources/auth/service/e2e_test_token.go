package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	jwtHelper "github.com/dcssoftware/bafoeg-manager/src/helper/jwt"
	jwtHelperModels "github.com/dcssoftware/bafoeg-manager/src/helper/jwt/models"
	mockProviderService "github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/e2e-mock"
	providerConst "github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
)

func (s *AuthService) CallbackE2ETestToken() (jwt, refreshToken string, err customerrors.ErrorInterface) {
	svc := mockProviderService.NewAuthService(s.userSvc)

	loggedinUserID, userErr := svc.MockCallbackFunction()
	if userErr != nil {
		return "", "", userErr
	}

	tx, txErr := s.storage.StartTransaction()
	if txErr != nil {
		return "", "", txErr
	}

	defer func() {
		_ = tx.Rollback()
	}()

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
