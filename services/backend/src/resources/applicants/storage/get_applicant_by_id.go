package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/storage/models"
	"github.com/go-sqlx/sqlx"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
)

func (s *ApplicantStorage) GetApplicantByID(tx *sqlx.Tx, applicantID string) (*serviceModels.ApplicantModel, customerrors.ErrorInterface) {

	sqlquery := `
			SELECT * FROM applicants_with_address_and_contact_data
			WHERE id = $1
	`

	var user models.ApplicantModel
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery, applicantID).StructScan(&user)
	} else {
		err = s.db.QueryRowx(sqlquery, applicantID).StructScan(&user)
	}

	if err != nil {
		data := customerrors.SQLData{
			"applicant": applicantID,
		}

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get user", sqlquery, data)
		}
	}

	return user.ToServiceModel(), nil
}
