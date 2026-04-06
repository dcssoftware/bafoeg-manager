package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *EakteStorage) GetEakteByID(tx *sqlx.Tx, id uuid.UUID) (*serviceModels.EakteModel, customerrors.ErrorInterface) {
	sqlquerybuilder := squirrel.
		Select(`*`).
		From("eakte_overview").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"id": id,
		}).OrderBy("id DESC")

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var row *sqlx.Row
	var err error

	if tx != nil {
		row = tx.QueryRowx(sqlquery, sqlArgs...)
	} else {
		row = s.db.QueryRowx(sqlquery, sqlArgs...)
	}

	var akte models.EakteModel
	err = row.StructScan(&akte)

	sqlErrorData := customerrors.SQLData{}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get eakte akte by id", sqlquery, sqlErrorData)
		}
	}

	return akte.ToServiceModel(), nil
}
