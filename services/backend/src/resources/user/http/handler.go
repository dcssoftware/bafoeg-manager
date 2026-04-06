package http

import (
	"github.com/dcssoftware/bafoeg-manager/src/resources/user/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
