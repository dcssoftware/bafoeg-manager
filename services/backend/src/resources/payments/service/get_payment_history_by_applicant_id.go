package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/payments/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *PaymentService) GetPaymentHistoryByApplicantID(tx *sqlx.Tx, page uint, applicantID string) ([]models.PaymentsModell, customerrors.ErrorInterface) {
	return s.storage.GetPaymentHistoryByApplicantID(tx, page, applicantID)
}
