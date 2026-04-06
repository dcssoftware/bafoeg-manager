package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModel "github.com/dcssoftware/bafoeg-manager/src/resources/organization/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *OrganizationStorage) GetAbteilungenCount(tx *sqlx.Tx, behördeID uuid.UUID) (uint, customerrors.ErrorInterface) {
	sqlquerybuilder := squirrel.Select(`COUNT(id) AS count`).From("responsible_behoerde_abteilung").PlaceholderFormat(squirrel.Dollar).Where(squirrel.Eq{
		"behoerde_id": behördeID,
	})

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
