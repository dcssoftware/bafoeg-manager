package storage

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/states"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsStorage) UpdateApplicationStatus(tx *sqlx.Tx, applicationID string, newStatus states.ApplicationState) customerrors.ErrorInterface {

	var err error
	sqlquery := `UPDATE applications SET status=(SELECT id FROM application_status WHERE identifier=$2) WHERE id=$1 RETURNING *`

	if tx != nil {
		_, err = tx.Queryx(
			sqlquery,
			applicationID,
			newStatus.ToString(),
		)
	} else {
		_, err = s.db.Queryx(
			sqlquery,
			applicationID,
			newStatus.ToString(),
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
