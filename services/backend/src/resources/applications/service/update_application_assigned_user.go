package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) UpdateApplicationAssignedUser(tx *sqlx.Tx, applicationID string, newAssignedUser string) customerrors.ErrorInterface {
	oldApplication, oldApplicationErr := s.GetApplicationByID(tx, applicationID)
	if oldApplicationErr != nil {
		return oldApplicationErr
	}

	s.storage.UpdateApplicationAssignedUser(tx, oldApplication.ID.String(), newAssignedUser)

	return nil
}
