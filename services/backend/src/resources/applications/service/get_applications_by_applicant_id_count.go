package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) GetApplicationsByApplicantIDCount(tx *sqlx.Tx, applicantID string) (uint, customerrors.ErrorInterface) {
	return s.storage.GetApplicationsByApplicantIDCount(tx, applicantID)
}
