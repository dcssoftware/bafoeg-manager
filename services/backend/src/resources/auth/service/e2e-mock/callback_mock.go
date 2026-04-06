package e2emock

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	providerConst "github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
	"github.com/go-sqlx/sqlx"
)

const (
	MockName     = "E2E Test User"
	MockUsername = "E2ETESTUSER"
)

func (s *AuthService) MockCallbackFunction() (string, customerrors.ErrorInterface) {

	user, userErr := s.userSvc.GetUserByUsername(nil, MockUsername)

	if userErr != nil && userErr.ErrorType() == customerrorconst.ERROR_IDENTIFIER_DATABASE_NOT_FOUND {
		userID, userErr := s.CreateUserWithLoginSession()
		if userErr != nil {
			return "", userErr
		}
		return userID, nil
	}

	return user.ID, nil
}

func (s *AuthService) CreateUserWithLoginSession() (
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
		MockName,
		MockUsername,
		"test@example.com",
		true,
	)
	if createUserErr != nil {
		_ = tx.Rollback()
		return "", createUserErr
	}

	loginUserID = createdUserID

	createdUserLoginIdentityErr := s.userSvc.CreateUserLoginIdentity(
		tx,
		createdUserID,
		providerConst.E2EMock,
		"",
	)
	if createdUserLoginIdentityErr != nil {
		_ = tx.Rollback()
		return "", createdUserLoginIdentityErr
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		errMsg := "(oauth callback e2e-mock) Failed to commit transaction (create user)"
		return "", customerrors.NewDatabaseError(commitErr, "", errMsg, "", nil)
	}

	return loginUserID, nil
}
