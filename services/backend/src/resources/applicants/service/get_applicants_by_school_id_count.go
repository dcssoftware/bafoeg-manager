package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicantService) GetApplicantsBySchoolIDCount(tx *sqlx.Tx, isActive bool, schoolID string) (uint, customerrors.ErrorInterface) {
	return s.storage.GetApplicantsBySchoolIDCount(tx, isActive, schoolID)
}
