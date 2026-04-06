package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsStorage) GetApplicationFileByFileID(tx *sqlx.Tx, fileID string) (*serviceModel.ApplicationFile, customerrors.ErrorInterface) {
	sqlquerybuilder := squirrel.
		Select(`*`).
		From("application_files_view").
		Where(squirrel.Eq{"id": fileID}).
		PlaceholderFormat(squirrel.Dollar).
		Limit(1)

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var row *sqlx.Row
	var err error

	if tx != nil {
		row = tx.QueryRowx(sqlquery, sqlArgs...)
	} else {
		row = s.db.QueryRowx(sqlquery, sqlArgs...)
	}

	var application models.ApplicationFile
	err = row.StructScan(&application)

	sqlErrorData := customerrors.SQLData{}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get application file by file id", sqlquery, sqlErrorData)
		}
	}

	return application.ToServiceFileModel(), nil
}
