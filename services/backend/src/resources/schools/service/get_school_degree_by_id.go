package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *SchoolService) GetSchoolDegreeByID(tx *sqlx.Tx, degreeID uuid.UUID) (*model.SchoolDegreeModel, customerrors.ErrorInterface) {
	return s.storage.GetSchoolDegreeByID(tx, degreeID)
}
