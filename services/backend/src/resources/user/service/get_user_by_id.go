package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *UserService) GetUserByID(tx *sqlx.Tx, userID string) (*models.UserModel, customerrors.ErrorInterface) {
	user, userErr := s.storage.GetByID(tx, userID)
	if userErr != nil {
		return nil, userErr
	}
	return user, nil
}
