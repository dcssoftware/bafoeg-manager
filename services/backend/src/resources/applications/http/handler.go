package http

import "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service"

type ApplicationsHandler struct {
	service *service.ApplicationsService
}

func NewApplicationsHandler(service *service.ApplicationsService) *ApplicationsHandler {
	return &ApplicationsHandler{
		service,
	}
}
