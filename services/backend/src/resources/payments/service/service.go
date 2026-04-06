package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/payments/service/models"
	"github.com/go-sqlx/sqlx"
)

type PaymentService struct {
	storage PaymentStorageInterface
}

func NewPaymentService(storage PaymentStorageInterface) *PaymentService {
	return &PaymentService{
		storage,
	}
}

type PaymentStorageInterface interface {
	GetPaymentHistoryByApplicantID(tx *sqlx.Tx, page uint, applicantID string) ([]models.PaymentsModell, customerrors.ErrorInterface)
	GetPaymentOutgoingTotalByApplicantID(tx *sqlx.Tx, applicantID string) (float64, customerrors.ErrorInterface)
}
