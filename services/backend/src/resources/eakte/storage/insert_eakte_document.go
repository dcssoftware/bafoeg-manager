package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *EakteStorage) InsertEakteDocument(tx *sqlx.Tx, vorgangID, fileID uuid.UUID, sourceIdentidier string, isXdomeaZipFile bool) (uuid.UUID, customerrors.ErrorInterface) {
	sqlquery := `INSERT INTO eakte_import_dokument (source, vorgang_id, file_id, source_xdomea_file) VALUES ((SELECT id FROM eakte_akte_source WHERE identifier = $1), $2, $3, $4) RETURNING id`

	var row *sqlx.Row
	var err error

	if tx != nil {
		row = tx.QueryRowx(sqlquery, sourceIdentidier, vorgangID, fileID, isXdomeaZipFile)
	} else {
		row = s.db.QueryRowx(sqlquery, sourceIdentidier, vorgangID, fileID, isXdomeaZipFile)
	}

	var vorgang models.IDModel
	err = row.StructScan(&vorgang)

	sqlErrorData := customerrors.SQLData{
		"vorgang_id":         vorgangID,
		"file_id":            fileID,
		"source_identifier":  sourceIdentidier,
		"is_xdomea_zip_file": isXdomeaZipFile,
	}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot insert eakte dokument", sqlquery, sqlErrorData)
		}
	}

	return vorgang.ID, nil
}
