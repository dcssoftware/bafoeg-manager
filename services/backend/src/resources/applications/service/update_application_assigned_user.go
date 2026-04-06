package service

import (
	"fmt"

	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) UpdateApplicationAssignedUser(tx *sqlx.Tx, applicationID string, newAssignedUser string) customerrors.ErrorInterface {
	oldApplication, oldApplicationErr := s.GetApplicationByID(tx, applicationID)
	if oldApplicationErr != nil {
		fmt.Println("Error: 1")
		return oldApplicationErr
	}

	s.storage.UpdateApplicationAssignedUser(tx, oldApplication.ID.String(), newAssignedUser)

	return nil
}
