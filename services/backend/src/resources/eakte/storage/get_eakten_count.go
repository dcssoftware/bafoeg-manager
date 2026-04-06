package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *EakteStorage) GetEaktenCount(tx *sqlx.Tx) (uint, customerrors.ErrorInterface) {
	sqlquery := `SELECT COUNT(id) AS "count" FROM eakte_overview`

	var model models.CountModel
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery).StructScan(&model)
	} else {
		err = s.db.QueryRowx(sqlquery).StructScan(&model)
	}

	if err != nil {
		data := customerrors.SQLData{}

		switch err {

		case sql.ErrNoRows:
			return 0, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return 0, customerrors.NewDatabaseError(err, "", "Cannot get eakten count", sqlquery, data)
		}
	}

	return model.Count, nil
}
