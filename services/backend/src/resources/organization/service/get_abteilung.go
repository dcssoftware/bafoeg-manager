package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/organization/service/models"
	"github.com/go-sqlx/sqlx"
	"github.com/google/uuid"
)

func (s *OrganizationService) GetAbteilungen(tx *sqlx.Tx, behördeID uuid.UUID, page uint) (uint, []models.Abteilung, customerrors.ErrorInterface) {

	count, countErr := s.storage.GetAbteilungenCount(tx, behördeID)
	if countErr != nil {
		return 0, nil, countErr
	}

	abteilungen, abteilungenErr := s.storage.GetAbteilungen(tx, behördeID, page)
	if abteilungenErr != nil {
		return 0, nil, abteilungenErr
	}

	return count, abteilungen, nil
}
