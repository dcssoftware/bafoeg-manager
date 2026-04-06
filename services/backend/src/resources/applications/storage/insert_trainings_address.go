package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsStorage) InsertTrainingsAddress(tx *sqlx.Tx, servideModel *serviceModels.ApplicationRevisionTrainingsdata) (uuid.UUID, customerrors.ErrorInterface) {
	var result storageModels.IDModel
	var row *sqlx.Row

	sqlquery := `INSERT INTO applicant_training_address (street, house_number, zip_code, city, country) VALUES ($1,$2,$3,$4,$5) RETURNING id`

	if tx != nil {
		row = tx.QueryRowx(
			sqlquery,
			servideModel.Street,
			servideModel.HouseNumber,
			servideModel.ZipCode,
			servideModel.City,
			servideModel.Country,
		)
	} else {
		row = s.db.QueryRowx(
			sqlquery,
			servideModel.Street,
			servideModel.HouseNumber,
			servideModel.ZipCode,
			servideModel.City,
			servideModel.Country,
		)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot insert application revision trainings address", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil
}
