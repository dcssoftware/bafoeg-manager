package http

import "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service"

type EakteHandler struct {
	service *service.EakteService
}

func NewEakteHandler(service *service.EakteService) *EakteHandler {
	return &EakteHandler{
		service: service,
	}
}
