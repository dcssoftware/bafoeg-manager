package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsStorage) InsertApplicationRevision(tx *sqlx.Tx, applicationID string, header, description string, baseRevision *serviceModels.ApplicationRevisionDataModel, newRevisionData *serviceModels.ApplicationRevisionDataModel) (uuid.UUID, customerrors.ErrorInterface) {

	var result storageModels.IDModel
	var row *sqlx.Row

	// Get training address ID and handle empty values
	trainingsAddressID := serviceModels.GetTrainingsAddressID(baseRevision, newRevisionData)
	var trainingsAddressParam *string
	if trainingsAddressID == nil || *trainingsAddressID == "" {
		trainingsAddressParam = nil
	} else {
		trainingsAddressParam = trainingsAddressID
	}

	sqlquery := `
		INSERT INTO application_revisions (
			message_header,
			message_description,
			application_id,
			trainings_address_id
		) VALUES
		 ($1,$2,$3,$4)
		 RETURNING id
	`

	if tx != nil {
		row = tx.QueryRowx(
			sqlquery,
			header,
			description,
			applicationID,
			trainingsAddressParam,
		)
	} else {
		row = s.db.QueryRowx(
			sqlquery,
			header,
			description,
			applicationID,
			trainingsAddressParam,
		)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot insert application file", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil
}
