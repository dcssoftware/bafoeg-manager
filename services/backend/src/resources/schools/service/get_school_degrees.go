package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/schools/service/model"
	"github.com/go-sqlx/sqlx"
)

func (s *SchoolService) GetSchoolDegree(tx *sqlx.Tx, page uint) ([]model.SchoolDegreeModel, customerrors.ErrorInterface) {
	return s.storage.GetSchoolDegree(tx, page)
}
