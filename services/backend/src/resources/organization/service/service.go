package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/organization/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

type OrganizationService struct {
	storage OrganizationStorage
}

func NewOrganizationService(storage OrganizationStorage) *OrganizationService {
	return &OrganizationService{storage: storage}
}

type OrganizationStorage interface {
	GetAbteilungen(tx *sqlx.Tx, behördeID uuid.UUID, page uint) ([]models.Abteilung, customerrors.ErrorInterface)
	GetAbteilungenCount(tx *sqlx.Tx, behördeID uuid.UUID) (uint, customerrors.ErrorInterface)

	GetBehörden(tx *sqlx.Tx, regionID uuid.UUID, page uint) ([]models.Behörde, customerrors.ErrorInterface)
	GetBehördenCount(tx *sqlx.Tx, regionID uuid.UUID) (uint, customerrors.ErrorInterface)

	GetRegions(tx *sqlx.Tx, page uint) ([]models.Region, customerrors.ErrorInterface)
	GetRegionsCount(tx *sqlx.Tx) (uint, customerrors.ErrorInterface)
}
