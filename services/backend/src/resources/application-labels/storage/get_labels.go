package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/application-labels/storage/models"
	"github.com/go-sqlx/sqlx"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/application-labels/service/models"
)

func (s *ApplicationLabelsStorage) GetLabels(tx *sqlx.Tx, page uint) ([]serviceModel.ApplicationLabel, customerrors.ErrorInterface) {
	var limit uint = configuration.Webserver.Display.MaxResponseEntityCount
	offset := (limit * page) - limit

	sqlquerybuilder := squirrel.
		Select(`*`).
		From("application_labels_with_color").
		PlaceholderFormat(squirrel.Dollar).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var rows *sqlx.Rows
	var applicationLabels []serviceModel.ApplicationLabel
	var err error

	if tx != nil {
		rows, err = tx.Queryx(sqlquery, sqlArgs...)
	} else {
		rows, err = s.db.Queryx(sqlquery, sqlArgs...)
	}

	sqlErrorData := customerrors.SQLData{}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get application labels", sqlquery, sqlErrorData)
		}
	}

	defer rows.Close()

	for rows.Next() {
		var application models.ApplicationLabel
		if err := rows.StructScan(&application); err != nil {
			return nil, customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
		}

		applicationLabels = append(applicationLabels, *application.ToApplicationLabelServiceModel())
	}

	return applicationLabels, nil
}
