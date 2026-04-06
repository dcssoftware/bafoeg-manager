package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"

	storageModel "github.com/dcssoftware/bafoeg-manager/src/resources/organization/storage/models"
)

func (s *OrganizationStorage) GetRegionsCount(tx *sqlx.Tx) (uint, customerrors.ErrorInterface) {
	sqlquerybuilder := squirrel.Select(`COUNT(id) AS count`).From("responsible_region").PlaceholderFormat(squirrel.Dollar)

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
		data := customerrors.SQLData{}

		switch err {

		case sql.ErrNoRows:
			return 0, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return 0, customerrors.NewDatabaseError(err, "", "Cannot get regions count", sqlquery, data)
		}
	}

	return model.Count, nil
}
