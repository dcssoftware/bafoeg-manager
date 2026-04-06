package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/storage/model"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *SchoolStorage) GetSchoolDegreeByID(tx *sqlx.Tx, degreeID uuid.UUID) (*serviceModel.SchoolDegreeModel, customerrors.ErrorInterface) {
	sqlquery := `
		SELECT * FROM school_degrees
		WHERE id = $1
		LIMIT 1
		`

	var row *sqlx.Row
	var degree model.SchoolDegreeModel
	var err error

	if tx != nil {
		row = tx.QueryRowx(sqlquery, degreeID.String())
	} else {
		row = s.db.QueryRowx(sqlquery, degreeID.String())
	}

	err = row.StructScan(&degree)

	sqlErrorData := customerrors.SQLData{}

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get school by id", sqlquery, sqlErrorData)
		}
	}

	degreeServiceModel := (&degree).ToServiceModel()
	return &degreeServiceModel, nil
}
