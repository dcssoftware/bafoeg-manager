package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *SchoolService) GetSchoolDegreeBySchoolIDCount(tx *sqlx.Tx, schoolID string) (uint, customerrors.ErrorInterface) {
	return s.storage.GetSchoolDegreeBySchoolIDCount(tx, schoolID)
}
