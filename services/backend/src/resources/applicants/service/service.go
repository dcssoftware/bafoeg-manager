package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/google/uuid"

	"github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service/models"
	paymentsService "github.com/dcssoftware/bafoeg-manager/src/resources/payments/service"
	"github.com/go-sqlx/sqlx"
)

type ApplicantService struct {
	storage     ApplicantStorageInterface
	paymentsSvc *paymentsService.PaymentService
}

func NewApplicantService(storage ApplicantStorageInterface, paymentsSvc *paymentsService.PaymentService) *ApplicantService {
	return &ApplicantService{
		storage,
		paymentsSvc,
	}
}

type ApplicantStorageInterface interface {
	GetApplicantByID(tx *sqlx.Tx, applicantID string) (*serviceModels.ApplicantModel, customerrors.ErrorInterface)
	GetApplicantsLastTrainingsAddressByApplicantID(tx *sqlx.Tx, applicantID string) (*serviceModels.ApplicantTrainingsAddressModel, customerrors.ErrorInterface)
	GetApplicants(tx *sqlx.Tx, page uint, filter string) ([]serviceModels.ApplicantModel, customerrors.ErrorInterface)
	GetApplicantsCount(tx *sqlx.Tx, filter string) (uint, customerrors.ErrorInterface)
	GetApplicantsBySchoolID(tx *sqlx.Tx, page uint, isActive bool, schoolID string) ([]serviceModels.ApplicantBySchoolModel, customerrors.ErrorInterface)
	GetApplicantsBySchoolIDCount(tx *sqlx.Tx, isActive bool, schoolID string) (uint, customerrors.ErrorInterface)

	InsertApplicant(tx *sqlx.Tx, model models.ApplicantInsertModel) (uuid.UUID, customerrors.ErrorInterface)
	InsertApplicantContactInformation(tx *sqlx.Tx, model models.ApplicantContactModel) (uuid.UUID, customerrors.ErrorInterface)
	InsertApplicantBankAccount(tx *sqlx.Tx, model models.ApplicantBankAccountModel) (uuid.UUID, customerrors.ErrorInterface)
	InsertApplicantAddress(tx *sqlx.Tx, model models.ApplicantAddressModel) (uuid.UUID, customerrors.ErrorInterface)
}
