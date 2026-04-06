package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *SchoolService) GetSchoolsCount(tx *sqlx.Tx, filter string) (uint, customerrors.ErrorInterface) {
	return s.storage.GetSchoolsCount(tx, filter)
}
