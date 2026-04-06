package service

import "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors"

func (s *MiddlewareService) GetUserPermissionsByID(userID string) ([]string, customerrors.ErrorInterface) {
	return s.userSvc.GetUserPermissionsByID(nil, userID)
}
