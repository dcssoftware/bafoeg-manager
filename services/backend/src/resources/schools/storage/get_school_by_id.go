package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/storage/model"
	"github.com/go-sqlx/sqlx"
)

func (s *SchoolStorage) GetSchoolByID(tx *sqlx.Tx, schoolID string) (*serviceModel.SchoolModel, customerrors.ErrorInterface) {
	sqlquery := `
		SELECT 
			schools.id, 
			jsonb_agg(school_degrees) AS "school_degree",
			schools.name, 
			schools.street, 
			schools.house_number, 
			schools.city, 
			schools.zip_code, 
			schools.country,
			school_types.name AS "school_type_name",
			school_types.identifier AS "school_type_identifier"
		FROM schools

		INNER JOIN school_types
			ON schools.school_type_id = school_types.id 

		LEFT JOIN school_degrees
			ON school_degrees.school_id = schools.id

		WHERE schools.id = $1

		GROUP BY schools.id, school_types.id

		LIMIT 1
		`

	var row *sqlx.Row
	var school model.SchoolModel
	var err error

	if tx != nil {
		row = tx.QueryRowx(sqlquery, schoolID)
	} else {
		row = s.db.QueryRowx(sqlquery, schoolID)
	}

	err = row.StructScan(&school)

	sqlErrorData := customerrors.SQLData{}

	if err != nil {

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get school by id", sqlquery, sqlErrorData)
		}
	}

	return school.ToServiceModel(), nil
}
