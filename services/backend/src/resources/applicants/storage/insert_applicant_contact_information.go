package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicantStorage) InsertApplicantContactInformation(tx *sqlx.Tx, model models.ApplicantContactModel) (uuid.UUID, customerrors.ErrorInterface) {
	var result storageModels.IDModel
	var row *sqlx.Row

	sqlquery := `INSERT INTO applicants_contact_data (phone, email) VALUES ( $1,$2 ) RETURNING id`

	if tx != nil {
		row = tx.QueryRowx(
			sqlquery,
			model.Phone,
			model.Email,
		)
	} else {
		row = s.db.QueryRowx(
			sqlquery,
			model.Phone,
			model.Email,
		)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot insert application contact information", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil
}
