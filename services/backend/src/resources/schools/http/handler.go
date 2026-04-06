package http

import "github.com/dcssoftware/bafoeg-manager/src/resources/schools/service"

type SchoolHandler struct {
	service *service.SchoolService
}

func NewSchoolHandler(service *service.SchoolService) *SchoolHandler {
	return &SchoolHandler{
		service,
	}
}
