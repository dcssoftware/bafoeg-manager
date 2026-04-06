package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicantService) GetApplicants(tx *sqlx.Tx, page uint, filter string) ([]models.ApplicantModel, customerrors.ErrorInterface) {
	return s.storage.GetApplicants(tx, page, filter)
}
