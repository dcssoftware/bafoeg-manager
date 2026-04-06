package storage

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *UserStore) CreateLoginIdentity(
	tx *sqlx.Tx,
	createdUserID string,
	authProvider string,
	authProviderID string,
) customerrors.ErrorInterface {

	sqlquery := "INSERT INTO public.user_identities (provider, provider_user_id, user_id) VALUES ($1, $2, $3)"

	_, err := tx.Exec(sqlquery, authProvider, authProviderID, createdUserID)
	if err != nil {
		data := customerrors.SQLData{
			"authProvider":   authProvider,
			"authProviderID": authProviderID,
			"userID":         createdUserID,
		}

		return customerrors.NewDatabaseError(err, "", "", sqlquery, data)
	}
	return nil
}
