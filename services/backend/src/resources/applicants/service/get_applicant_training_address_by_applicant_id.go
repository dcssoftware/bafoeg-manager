package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicantService) GetApplicantTrainingsAddressByApplicantID(tx *sqlx.Tx, applicantID string) (*models.ApplicantTrainingsAddressModel, customerrors.ErrorInterface) {
	trainingsAddress, trainingsAddressErr := s.storage.GetApplicantsLastTrainingsAddressByApplicantID(tx, applicantID)
	if trainingsAddressErr != nil {
		return nil, trainingsAddressErr
	}

	return trainingsAddress, nil
}
