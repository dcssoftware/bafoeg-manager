package storage

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsStorage) UpdateApplicationAssignedSchoolDegree(tx *sqlx.Tx, applicationID uuid.UUID, newSchoolID uuid.UUID) customerrors.ErrorInterface {
	var err error
	stmt := `UPDATE applications SET school_degree_id = $1 WHERE id = $2`

	if tx != nil {
		_, err = tx.Exec(
			stmt,
			newSchoolID,
			applicationID,
		)
	} else {
		_, err = s.db.Exec(
			stmt,
			newSchoolID,
			applicationID,
		)
	}

	if err != nil {

		sqlData := customerrors.SQLData{
			"application_id":         applicationID,
			"new_assigned_school_id": newSchoolID,
		}

		return customerrors.NewDatabaseError(err, "", "cannot update school degree assignment for application", stmt, sqlData)
	}

	return nil
}
