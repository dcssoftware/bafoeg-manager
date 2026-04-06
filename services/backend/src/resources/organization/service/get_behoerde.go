package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/organization/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *OrganizationService) GetBehörden(tx *sqlx.Tx, regionID uuid.UUID, page uint) (uint, []models.Behörde, customerrors.ErrorInterface) {
	behördenCount, behördenCountErr := s.storage.GetBehördenCount(tx, regionID)
	if behördenCountErr != nil {
		return 0, nil, behördenCountErr
	}

	behörden, behördenErr := s.storage.GetBehörden(tx, regionID, page)
	if behördenErr != nil {
		return 0, nil, behördenErr
	}

	return behördenCount, behörden, nil
}
