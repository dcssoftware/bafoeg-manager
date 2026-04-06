package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicantService) GetApplicantsBySchoolID(tx *sqlx.Tx, page uint, isActive bool, schoolID string) ([]models.ApplicantBySchoolModel, customerrors.ErrorInterface) {
	return s.storage.GetApplicantsBySchoolID(tx, page, isActive, schoolID)
}
