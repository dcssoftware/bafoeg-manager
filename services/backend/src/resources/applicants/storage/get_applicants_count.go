package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicantStorage) GetApplicantsCount(tx *sqlx.Tx, filter string) (uint, customerrors.ErrorInterface) {
	// sqlquery := `SELECT COUNT(id) AS "count" FROM applications WHERE assigned_user_id=$1`
	sqlquerybuilder := squirrel.Select(`COUNT(id) AS count`).
		From("applicants_with_address_and_contact_data").
		PlaceholderFormat(squirrel.Dollar)

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
