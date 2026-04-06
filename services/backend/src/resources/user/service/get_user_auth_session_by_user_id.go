package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	serviceModel "github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
)

func (s *UserService) GetUserAuthSessionByUserID(userID string) (*serviceModel.SessionModel, customerrors.ErrorInterface) {
	return s.storage.GetUserAuthSessionByUserID(nil, userID)
}
