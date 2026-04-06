package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/go-sqlx/sqlx"
)

func (s *UserService) GetUsersCount(tx *sqlx.Tx) (uint, customerrors.ErrorInterface) {
	return s.storage.GetUsersCount(tx)
}
