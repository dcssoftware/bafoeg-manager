package storage

import (
	"database/sql"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applications/service/models"
	storageModels "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicationsStorage) InsertApplicationFromEakte(tx *sqlx.Tx, model models.ApplicationFromEakteInsertModel) (uuid.UUID, customerrors.ErrorInterface) {

	var result storageModels.IDModel
	var row *sqlx.Row

	sqlquery := `
		INSERT INTO applications (
			applicant_id, 
			class_level, 
			status, 
			assigned_user_id, 
			school_degree_id, 
			application_validity_starts, 
			application_validity_ends
		) VALUES 
		 ($1,$2,
		 (SELECT id FROM application_status WHERE identifier = 'IN_PROGRESS'),
		NULL,
		 $3,$4,$5)
		 RETURNING id
	`

	if tx != nil {
		row = tx.QueryRowx(sqlquery,
			model.ApplicantID,
			model.ClassLevel,
			model.SchoolWithDegreeID,
			model.EducationStart,
			model.EducationEnd)
	} else {
		row = s.db.QueryRowx(sqlquery,
			model.ApplicantID,
			model.ClassLevel,
			model.SchoolWithDegreeID,
			model.EducationStart,
			model.EducationEnd)
	}

	err := row.StructScan(&result)

	sqlErrorData := customerrors.SQLData{}
	if err != nil {
		switch err {

		case sql.ErrNoRows:
			return uuid.Nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, sqlErrorData)

		default:
			return uuid.Nil, customerrors.NewDatabaseError(err, "", "Cannot get application by id", sqlquery, sqlErrorData)
		}
	}

	return result.ID, nil
}
