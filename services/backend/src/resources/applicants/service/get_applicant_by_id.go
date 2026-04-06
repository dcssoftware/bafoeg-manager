package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicantService) GetApplicantByID(tx *sqlx.Tx, applicantID string) (*models.ApplicantModel, customerrors.ErrorInterface) {
	applicantModel, applicantModellErr := s.storage.GetApplicantByID(tx, applicantID)
	if applicantModellErr != nil {
		return nil, applicantModellErr
	}

	trainingsAddress, trainingsAddressErr := s.GetApplicantTrainingsAddressByApplicantID(tx, applicantID)
	if trainingsAddressErr == nil {
		applicantModel.TrainingsAddress = trainingsAddress
	}

	outgoingBalance, outgoingBalanceErr := s.paymentsSvc.GetOutgoingPaymentsBalance(tx, applicantID)
	if outgoingBalanceErr == nil {
		applicantModel.BalanceOutgoing = &outgoingBalance
	}

	return applicantModel, nil
}
