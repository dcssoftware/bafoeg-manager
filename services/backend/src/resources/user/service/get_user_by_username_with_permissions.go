package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
	"github.com/go-sqlx/sqlx"
)

func (s *UserService) GetUserByUsernameWithPermissions(
	tx *sqlx.Tx,
	username string,
) (
	*models.UserWithPermissionsModel,
	customerrors.ErrorInterface,
) {

	user, userErr := s.GetUserByUsername(tx, username)
	if userErr != nil {
		return nil, userErr
	}

	permissions, permissionsErr := s.GetUserPermissionsByID(tx, user.ID)
	if permissionsErr != nil {
		return nil, permissionsErr
	}

	return models.NewUserWithPermissionsModel(user, permissions), nil
}
