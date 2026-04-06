package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/http/models"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *EakteHandler) GetEakteByID(c fiber.Ctx) error {

	eakteIDString := c.Params("eakteID", "")
	if eakteIDString == "" || !uuidvalidator.ValidateUUID(eakteIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	akte, akteErr := h.service.GetEakteByID(nil, uuid.MustParse(eakteIDString))
	if akteErr != nil {
		status, message := akteErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.Status(http.StatusOK).JSON(models.ToEakteHttpModel(akte))
}
