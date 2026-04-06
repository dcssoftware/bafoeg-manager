package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/storage/models"
	"github.com/go-sqlx/sqlx"

	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
)

func (s *ApplicantStorage) GetApplicants(tx *sqlx.Tx, page uint, filter string) ([]serviceModels.ApplicantModel, customerrors.ErrorInterface) {
	var limit uint = configuration.Webserver.Display.MaxResponseEntityCount
	offset := (limit * page) - limit

	sqlquerybuilder := squirrel.Select("*").
		From("applicants_with_address_and_contact_data").
		PlaceholderFormat(squirrel.Dollar).
		Limit(uint64(limit)).
		Offset(uint64(offset))

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
	var users []models.ApplicantModel
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
		var user models.ApplicantModel
		if err := rows.StructScan(&user); err != nil {
			customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
			panic(err)
		}
		users = append(users, user)
	}

	return models.ToApplicationServiceModels(users), nil
}
