package storage

import (
	"github.com/Masterminds/squirrel"
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/organization/service/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/organization/storage/models"
	"github.com/go-sqlx/sqlx"
)

func (s *OrganizationStorage) GetRegions(tx *sqlx.Tx, page uint) ([]serviceModel.Region, customerrors.ErrorInterface) {

	var limit uint = configuration.Webserver.Display.MaxResponseEntityCount
	offset := (limit * page) - limit

	sqlquerybuilder := squirrel.
		Select(`*`).
		From("responsible_region").
		PlaceholderFormat(squirrel.Dollar).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	sqlquery, sqlArgs, sqlErr := sqlquerybuilder.ToSql()
	if sqlErr != nil {
		return nil, customerrors.NewDatabaseError(sqlErr, "", "cannot build sql query", "", nil)
	}

	var rows *sqlx.Rows
	var regions []serviceModel.Region
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
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get regions", sqlquery, sqlErrorData)
		}
	}

	defer rows.Close()

	for rows.Next() {
		var region models.Region
		if err := rows.StructScan(&region); err != nil {
			return nil, customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
		}

		regions = append(regions, *region.ToServiceModel())
	}

	return regions, nil
}
