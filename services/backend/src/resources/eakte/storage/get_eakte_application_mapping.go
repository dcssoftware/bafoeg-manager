package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *EakteStorage) GetEakteApplicationMapping(tx *sqlx.Tx, eakteAkteID string) (*serviceModels.EaktenApplicationMappingModel, customerrors.ErrorInterface) {

	squirrelbuilder := squirrel.
		Select("*").
		From("eakte_application_mappings_overview").
		Where(squirrel.Eq{
			"eakte_akte_id": eakteAkteID,
		}).
		PlaceholderFormat(squirrel.Dollar).
		Limit(1)

	sqlquery, sqlArgs, sqlErr := squirrelbuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query to query eakten application mapping", "", nil)
	}

	var row *sqlx.Row
	var err error

	if tx != nil {
		row = tx.QueryRowx(sqlquery, sqlArgs...)
	} else {
		row = s.db.QueryRowx(sqlquery, sqlArgs...)
	}

	var mapping models.EaktenApplicationMappingModel
	err = row.StructScan(&mapping)

	sqlErrorData := customerrors.SQLData{
		"eakte_akte_id": eakteAkteID,
	}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get application by id", sqlquery, sqlErrorData)
		}
	}

	return mapping.ToServiceModel(), nil
}
