package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *EakteStorage) InsertEakteVorgang(tx *sqlx.Tx, akteID uuid.UUID, vorgangszeichen string) (uuid.UUID, customerrors.ErrorInterface) {
	sqlquery := `INSERT INTO eakte_import_vorgang (akte_id, vorgangszeichen) VALUES ($1, $2) RETURNING id`

	var row *sqlx.Row
	var err error

	if tx != nil {
		row = tx.QueryRowx(sqlquery, akteID, vorgangszeichen)
	} else {
		row = s.db.QueryRowx(sqlquery, akteID, vorgangszeichen)
	}

	var vorgang models.IDModel
	err = row.StructScan(&vorgang)

	sqlErrorData := customerrors.SQLData{
		"akte_id":         akteID,
		"vorgangszeichen": vorgangszeichen,
	}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot insert eakte vorgang", sqlquery, sqlErrorData)
		}
	}

	return vorgang.ID, nil
}
