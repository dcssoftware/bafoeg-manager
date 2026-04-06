package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *EakteStorage) GetFilesByVorgangsIDCount(tx *sqlx.Tx, vorgangID string) (uint, customerrors.ErrorInterface) {

	sqlquerybuilder := squirrel.Select(`COUNT(id) AS count`).
		From("eakte_files_overview").
		PlaceholderFormat(squirrel.Dollar).
		Where(
			squirrel.And{
				squirrel.Eq{
					"vorgang->>'id'": vorgangID,
				},
				squirrel.Eq{
					"source_xdomea_file": false,
				},
			},
		)

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
			return 0, customerrors.NewDatabaseError(err, "", "Cannot get eakten documents count", sqlquery, data)
		}
	}

	return model.Count, nil
}
