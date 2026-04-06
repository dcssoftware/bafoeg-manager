package service

import (
	"context"
	"encoding/json"
	"io"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	models "github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/authentik/models"
	providerConst "github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
	"github.com/go-sqlx/sqlx"
)

func (s *AuthService) AuthentikCallbackFunction(authCode, state string) (string, customerrors.ErrorInterface) {

	authToken, err := s.authentikConfig.Exchange(context.Background(), authCode)
	if err != nil {
		errMsg := "(oauth callback authentik) Failed to exchange auth code with authentik provider"
		return "", customerrors.NewInternalServerError(err, "", errMsg)
	}

	client := s.authentikConfig.Client(context.Background(), authToken)
	resp, err := client.Get(s.authentikOauthUserInfoEP)
	if err != nil {
		errMsg := "(oauth callback authentik) Failed to get user info from authentik provider"
		return "", customerrors.NewInternalServerError(err, "", errMsg)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := "(oauth callback authentik) Failed to read authentik oauth user response body"
		return "", customerrors.NewInternalServerError(err, "", errMsg)
	}

	var result *models.AuthentikUserInfo
	if err := json.Unmarshal(data, &result); err != nil {
		errMsg := "(oauth callback authentik) Failed to unmarshal authentik oauth user response body"
		return "", customerrors.NewInternalServerError(err, "", errMsg)
	}

	loginUserID, loginUserErr := s.userSvc.GetUserIDByLogin(providerConst.Authentik, result.ID)

	if loginUserErr != nil &&
		loginUserErr.ErrorType() != customerrorconst.ERROR_IDENTIFIER_DATABASE_NOT_FOUND {

		return "", loginUserErr

	} else if loginUserErr != nil {
		// Create new user if user does not exist or relationship cannot be established

		var newUserErr customerrors.ErrorInterface
		loginUserID, newUserErr = s.CreateUserWithLoginSession(result)
		if newUserErr != nil {
			return "", newUserErr
		}

	}

	return loginUserID, nil
}

func (s *AuthService) CreateUserWithLoginSession(
	result *models.AuthentikUserInfo,
) (
	string,
	customerrors.ErrorInterface,
) {
	var loginUserID string

	tx, txErr := s.userSvc.StartTransaction()
	if txErr != nil {
		return "", txErr
	}

	defer func(tx *sqlx.Tx) {
		_ = tx.Rollback()
	}(tx)

	createdUserID, createUserErr := s.userSvc.CreateUser(
		tx,
		result.Name,
		result.PreferredUsername,
		result.Email,
		result.EmailVerified,
	)
	if createUserErr != nil {
		_ = tx.Rollback()
		return "", createUserErr
	}

	loginUserID = createdUserID

	createdUserLoginIdentityErr := s.userSvc.CreateUserLoginIdentity(
		tx,
		createdUserID,
		providerConst.Authentik,
		result.ID,
	)
	if createdUserLoginIdentityErr != nil {
		_ = tx.Rollback()
		return "", createdUserLoginIdentityErr
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		errMsg := "(oauth callback authentik) Failed to commit transaction (create user)"
		return "", customerrors.NewDatabaseError(commitErr, "", errMsg, "", nil)
	}

	return loginUserID, nil
}
