package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/storage/models"
)

func (s *ApplicantStorage) GetApplicantsLastTrainingsAddressByApplicantID(tx *sqlx.Tx, applicantID string) (*serviceModels.ApplicantTrainingsAddressModel, customerrors.ErrorInterface) {
	sqlquery := `
			SELECT applicant_training_address.* FROM applications

			LEFT JOIN applicants ON applications.applicant_id = applicants.id
			LEFT JOIN application_revisions revisions ON revisions.application_id = applications.id
			LEFT JOIN applicant_training_address ON applicant_training_address.id = revisions.trainings_address_id


			WHERE applicants.id = $1
			ORDER BY applications.created, revisions DESC
	`

	var applicantTrainingsAddress models.ApplicantTrainingsAddressModel
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery, applicantID).StructScan(&applicantTrainingsAddress)
	} else {
		err = s.db.QueryRowx(sqlquery, applicantID).StructScan(&applicantTrainingsAddress)
	}

	if err != nil {
		data := customerrors.SQLData{
			"es_user_id": applicantID,
		}

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get applicants last training address", sqlquery, data)
		}
	}

	return applicantTrainingsAddress.ToServiceModel(), nil
}
