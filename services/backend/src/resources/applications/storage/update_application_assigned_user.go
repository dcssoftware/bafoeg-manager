package storage

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsStorage) UpdateApplicationAssignedUser(tx *sqlx.Tx, applicationID string, newAssignedUser string) customerrors.ErrorInterface {

	var err error
	sqlquery := `UPDATE applications SET assigned_user_id=(SELECT id FROM users WHERE id=$2) WHERE id=$1 RETURNING *`

	var queryParameter []any
	queryParameter = append(queryParameter, applicationID)

	// no set user is empty string so translate to nil for database
	if newAssignedUser == "" {
		queryParameter = append(queryParameter, nil)
	} else {
		queryParameter = append(queryParameter, newAssignedUser)
	}

	if tx != nil {
		_, err = tx.Queryx(
			sqlquery,
			queryParameter...,
		)
	} else {
		_, err = s.db.Queryx(
			sqlquery,
			queryParameter...,
		)
	}

	if err != nil {
		data := customerrors.SQLData{}

		switch err {

		default:
			return customerrors.NewDatabaseError(err, "", "Cannot get application count", sqlquery, data)
		}
	}

	return nil
}
