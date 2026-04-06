package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) GetApplicationRevisionsByApplicationIDCount(tx *sqlx.Tx, applicationID string) (uint, customerrors.ErrorInterface) {
	return s.storage.GetApplicationRevisionsByApplicationIDCount(tx, applicationID)
}
