package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *PaymentService) GetOutgoingPaymentsBalance(tx *sqlx.Tx, applicantID string) (float64, customerrors.ErrorInterface) {
	return s.storage.GetPaymentOutgoingTotalByApplicantID(tx, applicantID)
}
