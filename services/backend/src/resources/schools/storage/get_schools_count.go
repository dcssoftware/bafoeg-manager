package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/storage/model"
	"github.com/go-sqlx/sqlx"
)

func (s *SchoolStorage) GetSchoolsCount(tx *sqlx.Tx, filter string) (uint, customerrors.ErrorInterface) {

	sqlquerybuilder := squirrel.
		Select(`COUNT(id)`).
		From("school_overview").
		PlaceholderFormat(squirrel.Dollar)

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
		return 0, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var model model.CountModel
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
			return 0, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return 0, customerrors.NewDatabaseError(err, "", "Cannot get application count", sqlquery, data)
		}
	}

	return model.Count, nil
}
