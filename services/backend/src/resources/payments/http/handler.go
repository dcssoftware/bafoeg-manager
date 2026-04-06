package http

import "github.com/dcssoftware/bafoeg-manager/src/resources/payments/service"

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service,
	}
}
