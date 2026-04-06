package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicantStorage) GetApplicantsBySchoolID(tx *sqlx.Tx, page uint, isActive bool, schoolID string) ([]serviceModels.ApplicantBySchoolModel, customerrors.ErrorInterface) {
	var limit uint = configuration.Webserver.Display.MaxResponseEntityCount
	offset := (limit * page) - limit

	var sqlquerybuilder squirrel.SelectBuilder
	sqlquerybuilder = squirrel.
		Select("*").
		From("school_applicants_view").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"school_applicants_view.school_id": schoolID}).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	if isActive {
		sqlquerybuilder = sqlquerybuilder.
			Where("application_validity_starts <= now()").
			Where("application_validity_ends >= now()")
	}

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var rows *sqlx.Rows
	var applicants []models.ApplicantsBySchoolModel
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
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get applicants by school ID", sqlquery, sqlErrorData)
		}
	}

	defer rows.Close()

	for rows.Next() {
		var user models.ApplicantsBySchoolModel
		if err := rows.StructScan(&user); err != nil {
			customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
		}
		applicants = append(applicants, user)
	}

	return models.ToApplicantsBySchoolServiceModels(applicants), nil
}
