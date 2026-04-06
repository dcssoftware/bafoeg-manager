package http

import "github.com/dcssoftware/bafoeg-manager/src/resources/general/service"

type GeneralHandler struct {
	service *service.GeneralSvc
}

func NewGeneralHandler(service *service.GeneralSvc) *GeneralHandler {
	return &GeneralHandler{
		service,
	}
}
