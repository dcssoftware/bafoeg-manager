package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/go-sqlx/sqlx"
)

func (s *SchoolService) GetSchoolDegreeBySchoolID(tx *sqlx.Tx, page uint, schoolID string) ([]model.SchoolDegreeModel, customerrors.ErrorInterface) {
	return s.storage.GetSchoolDegreeBySchoolID(tx, page, schoolID)
}
