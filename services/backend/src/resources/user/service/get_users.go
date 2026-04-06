package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModels "github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *UserService) GetUsers(tx *sqlx.Tx) ([]serviceModels.UserModel, customerrors.ErrorInterface) {
	return s.storage.GetUsers(tx)
}
