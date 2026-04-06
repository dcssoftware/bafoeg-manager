package http

import (
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	uuidvalidator "github.com/dcssoftware/bafoeg-manager/src/helper/uuid-validator"
	"github.com/dcssoftware/bafoeg-manager/src/resources/eakte/http/models"
	"github.com/gofiber/fiber/v3"
)

func (h *EakteHandler) GetVorgänge(c fiber.Ctx) error {
	eakteIDString := c.Params("eakteID", "")
	if eakteIDString == "" || !uuidvalidator.ValidateUUID(eakteIDString) {
		return c.Status(http.StatusBadRequest).SendString(customerrorconst.BAD_REQUEST_ERROR_MESSAGE)
	}

	vorgänge, vorgängeErr := h.service.GetVorgängeByEaktenID(nil, eakteIDString)
	if vorgängeErr != nil {
		status, message := vorgängeErr.HTTPError()
		return c.Status(status).SendString(message)
	}

	return c.JSON(models.ToVorgängeHttpModel(vorgänge))
}
