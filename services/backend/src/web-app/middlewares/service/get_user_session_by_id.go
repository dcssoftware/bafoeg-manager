package service

import (
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"
	"github.com/dcssoftware/bafoeg-manager/src/resources/user/service/models"
)

func (s *MiddlewareService) GetUserSessionByID(sessionID string) (*models.SessionModel, customerrors.ErrorInterface) {
	return s.userSvc.GetUserAuthSessionByID(nil, sessionID)
}
