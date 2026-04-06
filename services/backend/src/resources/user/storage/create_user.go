package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *UserStore) Create(
	tx *sqlx.Tx,
	displayName string,
	username string,
	email string,
	emailVerified bool,
) (
	string,
	customerrors.ErrorInterface,
) {

	sqlquery := `
	INSERT INTO public.users 
		(display_name, username, email, email_verified) 
	VALUES 
		($1, $2, $3, $4) 
	RETURNING id`

	var userID string
	var err error
	if tx != nil {
		err = tx.Get(&userID, sqlquery, displayName, username, email, emailVerified)
	} else {
		err = s.db.Get(&userID, sqlquery, displayName, username, email, emailVerified)
	}

	if err != nil {

		data := customerrors.SQLData{
			"displayName":   displayName,
			"username":      username,
			"email":         email,
			"emailVerified": emailVerified,
		}

		switch err {

		case sql.ErrNoRows:
			return "", customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return "", customerrors.NewDatabaseError(err, "", "", sqlquery, data)
		}
	}

	return userID, nil
}
