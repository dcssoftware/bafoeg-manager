package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *UserService) CreateUser(
	tx *sqlx.Tx,
	name, username, email string,
	emailVerified bool,
) (
	string,
	customerrors.ErrorInterface,
) {

	var setTxToNotNil bool = false
	if tx == nil {

		var err customerrors.ErrorInterface
		tx, err = s.storage.StartTransaction()
		if err != nil {
			return "", err
		}

		defer tx.Rollback()
		setTxToNotNil = true

	}

	userCount, userCountErr := s.GetUsersCount(tx)
	if userCountErr != nil {
		return "", userCountErr
	}

	if userCount == 0 && configuration.ApplicationConfiguration.IsDevEnvironment {
		name = "Application Developer"
		username = "developer"
		email = "noreply@example.com"
		emailVerified = true
	}

	userID, userIDErr := s.storage.Create(tx, name, username, email, emailVerified)
	if userIDErr != nil {
		return "", userIDErr
	}

	// upload picture
	profilePictureErr := s.UploadProfilePictureByID(tx, userID, []byte{})
	if profilePictureErr != nil {
		return "", profilePictureErr
	}

	if setTxToNotNil {
		err := tx.Commit()
		if err != nil {
			return "", customerrors.NewInternalServerError(err, "", "could not create new user. Rollback transactions")
		}
	}

	return userID, nil
}
