package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *ApplicationsService) GetApplicationFilesByApplicationIDCount(tx *sqlx.Tx, applicationID string) (uint, customerrors.ErrorInterface) {
	return s.storage.GetApplicationFilesByApplicationIDCount(tx, applicationID)
}
