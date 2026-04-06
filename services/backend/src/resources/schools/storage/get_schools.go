package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/storage/model"
	"github.com/go-sqlx/sqlx"
)

func (s *SchoolStorage) GetSchools(tx *sqlx.Tx, page uint, filter string) ([]serviceModel.SchoolShortModel, customerrors.ErrorInterface) {
	var limit uint = configuration.Webserver.Display.MaxResponseEntityCount
	offset := (limit * page) - limit

	sqlquerybuilder := squirrel.
		Select(`*`).
		From("school_overview").
		PlaceholderFormat(squirrel.Dollar).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	if filter != "" {
		filter = "%" + filter + "%"
		sqlquerybuilder = sqlquerybuilder.Where(
			squirrel.Or{
				squirrel.Like{
					"id::text": filter,
				},
				squirrel.Like{
					"name": filter,
				},
				squirrel.Like{
					"street": filter,
				},
			},
		)
	}

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var rows *sqlx.Rows
	var schools []serviceModel.SchoolShortModel
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
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get school", sqlquery, sqlErrorData)
		}
	}

	defer rows.Close()

	for rows.Next() {
		var school model.SchoolShortModel
		if err := rows.StructScan(&school); err != nil {
			return nil, customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
		}

		schools = append(schools, *school.ToServiceShortModel())
	}

	return schools, nil
}
