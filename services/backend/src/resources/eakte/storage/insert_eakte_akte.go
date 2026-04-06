package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	stateModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/states/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *EakteStorage) InsertEakteAkte(tx *sqlx.Tx, aktenzeichen string, typIdentifier string, vertraulichkeit stateModels.VertraulichkeitEnum) (uuid.UUID, customerrors.ErrorInterface) {
	sqlquery := `INSERT INTO eakte_import_akte (aktenzeichen, typ, vertraulichkeit) VALUES ($1, (SELECT id FROM eakte_import_akte_type WHERE identifier = $2), $3) RETURNING id`

	var row *sqlx.Row
	var err error

	if tx != nil {
		row = tx.QueryRowx(sqlquery, aktenzeichen, typIdentifier, vertraulichkeit)
	} else {
		row = s.db.QueryRowx(sqlquery, aktenzeichen, typIdentifier, vertraulichkeit)
	}

	var akte models.IDModel
	err = row.StructScan(&akte)

	sqlErrorData := customerrors.SQLData{
		"aktenzeichen":    aktenzeichen,
		"typ_identifier":  typIdentifier,
		"vertraulichkeit": vertraulichkeit,
	}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot insert eakte akte", sqlquery, sqlErrorData)
		}
	}

	return akte.ID, nil
}
