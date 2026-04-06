package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	httpParams "github.com/dcssoftware/bafoeg-manager/src/helper/http-params"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/dcssoftware/bafoeg-manager/src/resources/payments/http/models"
	"github.com/gofiber/fiber/v3"
)

func (h *PaymentHandler) GetPaymentHistoryByApplicantID(c fiber.Ctx) error {
	queries := c.Queries()

	applicantIDString := c.Params("id", "")
	if applicantIDString == "" || !uuidvalidator.ValidateUUID(applicantIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	pageNumber, pageNumberErr := httpParams.GetParamsPageUint(queries)
	if pageNumberErr != nil {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	payments, paymentsErr := h.service.GetPaymentHistoryByApplicantID(nil, pageNumber, applicantIDString)
	if paymentsErr != nil {
		status, message := paymentsErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	paymentsJSON := models.ToHttpPaymentModels(payments)

	return c.Status(http.StatusOK).JSON(paymentsJSON)
}
