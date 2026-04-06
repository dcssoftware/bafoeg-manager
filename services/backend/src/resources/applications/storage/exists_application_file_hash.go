package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsStorage) ExistsUploadedApplicationFileHash(tx *sqlx.Tx, applicationID, filetype string, fileSize uint, fileHash string) (bool, customerrors.ErrorInterface) {

	sqlquerybuilder := squirrel.Select("COUNT(application_files.id) AS count").
		From("application_files").
		PlaceholderFormat(squirrel.Dollar).
		Join("files ON files.id = application_files.file_id").
		Where(
			squirrel.And{
				squirrel.Eq{"application_files.application_id": applicationID},
				squirrel.Eq{"files.file_size": fileSize},
				squirrel.Eq{"files.file_hash": fileHash},
			},
		)

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return false, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var model storageModel.CountModel
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery, sqlArgs...).StructScan(&model)
	} else {
		err = s.db.QueryRowx(sqlquery, sqlArgs...).StructScan(&model)
	}

	if err != nil {
		data := customerrors.SQLData{}

		switch err {

		case sql.ErrNoRows:
			return false, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return false, customerrors.NewDatabaseError(err, "", "Cannot get application count", sqlquery, data)
		}
	}

	return model.Count >= 1, nil
}
