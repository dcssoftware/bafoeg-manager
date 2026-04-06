package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *EakteStorage) GetFileByFileID(tx *sqlx.Tx, fileID uuid.UUID) (*serviceModels.EaktenFileModel, customerrors.ErrorInterface) {
	sqlquery := `SELECT * FROM eakte_import_dokument WHERE id = $1`

	var row *sqlx.Row
	var err error

	if tx != nil {
		row = tx.QueryRowx(sqlquery, fileID)
	} else {
		row = s.db.QueryRowx(sqlquery, fileID)
	}

	var eakteDocument models.EaktenFileModel
	err = row.StructScan(&eakteDocument)

	sqlErrorData := customerrors.SQLData{}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get eakten file by id", sqlquery, sqlErrorData)
		}
	}

	return eakteDocument.ToServiceModel(), nil
}
