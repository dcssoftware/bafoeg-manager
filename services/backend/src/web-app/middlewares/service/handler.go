package service

import userService "github.com/dcssoftware/bafoeg-manager/src/resources/user/service"

type MiddlewareService struct {
	userSvc *userService.UserService
}

func NewMiddlewareService(userSvc *userService.UserService) *MiddlewareService {
	return &MiddlewareService{
		userSvc: userSvc,
	}
}
