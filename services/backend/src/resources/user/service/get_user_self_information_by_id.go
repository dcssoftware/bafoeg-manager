package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
)

func (s *UserService) GetSelfInformationByID(
	userID string,
) (
	*serviceModel.UserSelfInformationModel,
	customerrors.ErrorInterface,
) {

	// get normal user entity
	user, userErr := s.storage.GetByID(nil, userID)
	if userErr != nil {
		return nil, userErr
	}

	permissions, permissionsErr := s.GetUserPermissionsByID(nil, userID)
	if permissionsErr != nil {
		return nil, permissionsErr
	}

	return serviceModel.NewUserSelfInformationModel(user, permissions), nil
}
