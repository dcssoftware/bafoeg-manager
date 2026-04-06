package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
)

func (s *ApplicationsStorage) GetApplicationByID(tx *sqlx.Tx, applicationID string) (*serviceModel.ApplicationModel, customerrors.ErrorInterface) {
	sqlquerybuilder := squirrel.
		Select(`*`).
		From("applications_view").
		Where(squirrel.Eq{"id": applicationID}).
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

	var application models.ApplicationModel
	err = row.StructScan(&application)

	sqlErrorData := customerrors.SQLData{}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get application by id", sqlquery, sqlErrorData)
		}
	}

	return application.ToServiceModel(), nil
}
