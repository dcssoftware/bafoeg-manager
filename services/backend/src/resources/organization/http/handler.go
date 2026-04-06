package http

import "github.com/dcssoftware/bafoeg-manager/src/resources/organization/service"

type OrganizationHandler struct {
	service *service.OrganizationService
}

func NewOrganizationHandler(service *service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{service: service}
}
