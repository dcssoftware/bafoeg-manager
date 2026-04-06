package storage

import (
	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/organization/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/organization/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *OrganizationStorage) GetBehörden(tx *sqlx.Tx, regionID uuid.UUID, page uint) ([]serviceModel.Behörde, customerrors.ErrorInterface) {

	var limit uint = configuration.Webserver.Display.MaxResponseEntityCount
	offset := (limit * page) - limit

	sqlquerybuilder := squirrel.
		Select(`*`).
		From("responsible_behoerde").
		PlaceholderFormat(squirrel.Dollar).
		Where(
			squirrel.Eq{"region_id": regionID},
		).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var rows *sqlx.Rows
	var regions []serviceModel.Behörde
	var err error

	if tx != nil {
		rows, err = tx.Queryx(sqlquery, sqlArgs...)
	} else {
		rows, err = s.db.Queryx(sqlquery, sqlArgs...)
	}

	sqlErrorData := customerrors.SQLData{}

	if err != nil {
		switch err {

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get behörden", sqlquery, sqlErrorData)
		}
	}

	defer rows.Close()

	for rows.Next() {
		var behörde models.Behörde
		if err := rows.StructScan(&behörde); err != nil {
			return nil, customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
		}

		regions = append(regions, *behörde.ToServiceModel())
	}

	return regions, nil
}
