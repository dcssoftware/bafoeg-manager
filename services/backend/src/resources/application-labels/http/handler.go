package http

import "github.com/dcssoftware/bafoeg-manager/src/resources/application-labels/service"

type ApplicationLabelsHandler struct {
	service *service.ApplicationLabelsService
}

func NewApplicationLabelsHandler(service *service.ApplicationLabelsService) *ApplicationLabelsHandler {
	return &ApplicationLabelsHandler{
		service,
	}
}
