package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModel "github.com/dcssoftware/bafoeg-manager/src/resources/schools/storage/model"
	"github.com/go-sqlx/sqlx"
)

func (s *SchoolStorage) GetSchoolDegreeBySchoolIDCount(tx *sqlx.Tx, schoolID string) (uint, customerrors.ErrorInterface) {

	sqlquerybuilder := squirrel.
		Select(`COUNT(id) AS count`).
		From("school_degrees").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"school_id": schoolID})

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return 0, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var model storageModel.CountModel
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery, sqlArgs...).StructScan(&model)
	} else {
		err = s.db.QueryRowx(sqlquery, sqlArgs...).StructScan(&model)
	}

	if err != nil {
		data := customerrors.SQLData{
			"school_id": schoolID,
		}

		switch err {

		case sql.ErrNoRows:
			return 0, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return 0, customerrors.NewDatabaseError(err, "", "Cannot get school degree count", sqlquery, data)
		}
	}

	return model.Count, nil
}
