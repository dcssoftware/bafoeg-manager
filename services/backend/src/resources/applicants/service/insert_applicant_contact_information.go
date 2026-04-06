package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *ApplicantService) InsertApplicantContactInformation(tx *sqlx.Tx, model models.ApplicantContactModel) (uuid.UUID, customerrors.ErrorInterface) {
	return s.storage.InsertApplicantContactInformation(tx, model)
}
