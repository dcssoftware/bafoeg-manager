package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicantStorage) InsertApplicantAddress(tx *sqlx.Tx, model models.ApplicantAddressModel) (uuid.UUID, customerrors.ErrorInterface) {
	var result storageModels.IDModel
	var row *sqlx.Row

	sqlquery := `
		INSERT INTO applicant_permanent_address (
			street,
			house_number,
			zip_code,
			city,
			country
		) VALUES
		 ($1,$2,$3,$4,$5)
		 RETURNING id
	`

	if tx != nil {
		row = tx.QueryRowx(
			sqlquery,
			model.Street,
			model.HouseNumber,
			model.ZipCode,
			model.City,
			model.Country,
		)
	} else {
		row = s.db.QueryRowx(
			sqlquery,
			model.Street,
			model.HouseNumber,
			model.ZipCode,
			model.City,
			model.Country,
		)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot insert application address", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil
}
