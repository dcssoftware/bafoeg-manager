package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/storage/model"
	"github.com/go-sqlx/sqlx"
)

func (s *SchoolStorage) GetSchoolDegree(tx *sqlx.Tx, page uint) ([]serviceModel.SchoolDegreeModel, customerrors.ErrorInterface) {
	sqlquery := `
		SELECT *
		FROM school_degrees
		OFFSET $1 
		LIMIT $2
		`

	var limit uint = configuration.Webserver.Display.MaxResponseEntityCount
	offset := (limit * page) - limit

	var rows *sqlx.Rows
	var schools []serviceModel.SchoolDegreeModel
	var err error

	if tx != nil {
		rows, err = tx.Queryx(sqlquery, offset, limit)
	} else {
		rows, err = s.db.Queryx(sqlquery, offset, limit)
	}

	sqlErrorData := customerrors.SQLData{}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get user", sqlquery, sqlErrorData)
		}
	}

	defer rows.Close()

	for rows.Next() {
		var degree model.SchoolDegreeModel
		if err := rows.StructScan(&degree); err != nil {
			return nil, customerrors.NewDatabaseError(err, "", "Could not read row from database", "sqlquery", sqlErrorData)
		}

		schools = append(schools, degree.ToServiceModel())
	}

	return schools, nil
}
