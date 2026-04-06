package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/organization/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *OrganizationService) GetRegions(tx *sqlx.Tx, page uint) (uint, []models.Region, customerrors.ErrorInterface) {
	count, countErr := s.storage.GetRegionsCount(tx)
	if countErr != nil {
		return 0, nil, countErr
	}

	regions, regionsErr := s.storage.GetRegions(tx, page)
	if regionsErr != nil {
		return 0, nil, regionsErr
	}

	return count, regions, nil
}
