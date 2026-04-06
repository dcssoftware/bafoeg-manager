package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicantService) InsertApplicantUnittest(tx *sqlx.Tx, model models.ApplicantInsertModel) (*models.ApplicantModel, customerrors.ErrorInterface) {
	insertedApplicantID, insertedApplicantErr := s.storage.InsertApplicant(tx, model)
	if insertedApplicantErr != nil {
		return nil, insertedApplicantErr
	}

	return s.GetApplicantByID(tx, insertedApplicantID.String())
}
