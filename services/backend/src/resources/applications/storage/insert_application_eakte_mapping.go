package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsStorage) InsertApplicationEakteMapping(tx *sqlx.Tx, applicationID, eakteAkteID uuid.UUID) (uuid.UUID, customerrors.ErrorInterface) {

	var result storageModels.IDModel
	var row *sqlx.Row

	sqlquerybuilder := squirrel.
		Insert("application_eakte_mapping").
		SetMap(map[string]any{
			"application_id": applicationID,
			"eakte_akte_id":  eakteAkteID,
		}).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar)

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return uuid.Nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	if tx != nil {
		row = tx.QueryRowx(sqlquery, sqlArgs...)
	} else {
		row = s.db.QueryRowx(sqlquery, sqlArgs...)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot insert application_eakte_mapping", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil
}
