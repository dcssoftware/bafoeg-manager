package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicantStorage) GetApplicantsBySchoolIDCount(tx *sqlx.Tx, isActive bool, schoolID string) (uint, customerrors.ErrorInterface) {
	sqlquerybuilder := squirrel.
		Select("COUNT(school_id)").
		From("school_applicants_view").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"school_applicants_view.school_id": schoolID})

	if isActive {
		sqlquerybuilder.
			Where(squirrel.Lt{"application_validity_starts": "now()"}).
			Where(squirrel.Gt{"application_validity_ends": "now()"})

	}

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return 0, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var model models.CountModel
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
