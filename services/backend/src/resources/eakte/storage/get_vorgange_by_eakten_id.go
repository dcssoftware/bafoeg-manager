package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *EakteStorage) GetVorgängeByEaktenID(tx *sqlx.Tx, id string) ([]serviceModels.VorgangModel, customerrors.ErrorInterface) {

	sqlquerybuilder := squirrel.
		Select(`*`).
		From("eakte_import_vorgang").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"akte_id": id,
		}).OrderBy("id DESC")

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var rows *sqlx.Rows
	var vorgänge []serviceModels.VorgangModel
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
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get vorgang (eakte)", sqlquery, sqlErrorData)
		}
	}

	defer rows.Close()

	for rows.Next() {
		var vorgang models.VorgangModel
		if err := rows.StructScan(&vorgang); err != nil {
			return nil, customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
		}

		vorgänge = append(vorgänge, *vorgang.ToServiceModel())
	}

	return vorgänge, nil
}
