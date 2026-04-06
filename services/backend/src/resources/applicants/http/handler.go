package http

import "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service"

type ApplicantHandler struct {
	service *service.ApplicantService
}

func NewApplicantHandler(service *service.ApplicantService) *ApplicantHandler {
	return &ApplicantHandler{
		service,
	}
}
