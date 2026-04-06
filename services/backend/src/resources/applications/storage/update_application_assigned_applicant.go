package storage

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsStorage) UpdateApplicationAssignedApplicant(tx *sqlx.Tx, applicationID uuid.UUID, newApplicantID uuid.UUID) customerrors.ErrorInterface {
	var err error
	stmt := `UPDATE applications SET applicant_id = $1 WHERE id = $2`

	if tx != nil {
		_, err = tx.Exec(
			stmt,
			newApplicantID,
			applicationID,
		)
	} else {
		_, err = s.db.Exec(
			stmt,
			newApplicantID,
			applicationID,
		)
	}

	if err != nil {

		sqlData := customerrors.SQLData{
			"application_id":            applicationID,
			"new_assigned_applicant_id": newApplicantID,
		}

		return customerrors.NewDatabaseError(err, "", "cannot update applicant assignment for application", stmt, sqlData)
	}

	return nil
}
