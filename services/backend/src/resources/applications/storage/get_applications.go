package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"

	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
)

func (s *ApplicationsStorage) GetApplications(tx *sqlx.Tx, page uint, userID string, filter string, filterOnlyInProgress bool) ([]serviceModel.ApplicationShortModel, customerrors.ErrorInterface) {
	var limit uint = configuration.Webserver.Display.MaxResponseEntityCount
	offset := (limit * page) - limit

	sqlquerybuilder := squirrel.
		Select(`*`).
		From("applications_overview").
		PlaceholderFormat(squirrel.Dollar).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	if userID != "" {
		sqlquerybuilder = sqlquerybuilder.Where(squirrel.Eq{"assigned_user_id": userID})
	}

	if filterOnlyInProgress {
		sqlquerybuilder = sqlquerybuilder.Where(
			squirrel.And{
				squirrel.NotEq{
					"status_identifier": "APPROVED",
				},
				squirrel.NotEq{
					"status_identifier": "DENIED",
				},
			},
		)
	}

	if filter != "" {
		sqlquerybuilder = sqlquerybuilder.Where(
			squirrel.Or{
				squirrel.Like{"id::text": "%" + filter + "%"},
				squirrel.Like{"applicant->>'firstname'": "%" + filter + "%"},
				squirrel.Like{"applicant->>'lastname'": "%" + filter + "%"},
			},
		)
	}

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var rows *sqlx.Rows
	var applications []serviceModel.ApplicationShortModel
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
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get user", sqlquery, sqlErrorData)
		}
	}

	defer rows.Close()

	for rows.Next() {
		var application models.ApplicationShortModel
		if err := rows.StructScan(&application); err != nil {
			return nil, customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
		}

		applications = append(applications, *application.ToServiceShortModel())
	}

	return applications, nil
}
