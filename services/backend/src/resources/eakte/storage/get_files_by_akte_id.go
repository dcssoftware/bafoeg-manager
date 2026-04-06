package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *EakteStorage) GetFilesByAkteID(tx *sqlx.Tx, akteID string) ([]serviceModels.DokumentModel, customerrors.ErrorInterface) {

	sqlquerybuilder := squirrel.
		Select(`*`).
		From("eakte_files_overview").
		PlaceholderFormat(squirrel.Dollar).
		Where(
			squirrel.And{
				squirrel.Eq{
					"akte->>'id'": akteID,
				},
				squirrel.Eq{
					"source_xdomea_file": false,
				},
			},
		).
		OrderBy("id DESC")

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var rows *sqlx.Rows
	var documents []models.DokumentModel
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
		var entry models.DokumentModel
		err := rows.StructScan(&entry)
		if err != nil {
			return nil, customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
		}
		documents = append(documents, entry)
	}

	return models.ToDokumentenServiceModel(documents), nil
}
