package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsStorage) GetApplicationRevisionsByApplicationID(tx *sqlx.Tx, page uint, applicantID string) ([]serviceModel.ApplicationRevisionShortModel, customerrors.ErrorInterface) {
	var limit uint = configuration.Webserver.Display.MaxResponseEntityCount
	offset := (limit * page) - limit

	sqlquerybuilder := squirrel.
		Select(`id, application_id, message_header, message_description, created`).
		From("application_revisions").
		PlaceholderFormat(squirrel.Dollar).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		OrderBy("id DESC").
		Where(squirrel.Eq{"application_id": applicantID})

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var rows *sqlx.Rows
	var applicationsRevision []serviceModel.ApplicationRevisionShortModel
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
		var application models.ApplicationRevisionShortModel
		if err := rows.StructScan(&application); err != nil {
			return nil, customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
		}

		applicationsRevision = append(applicationsRevision, *application.ToServiceShortModel())
	}

	if len(applicationsRevision) < 1 {
		return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)
	}

	return applicationsRevision, nil

}
