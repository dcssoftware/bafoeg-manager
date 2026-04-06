package http

import (
	"encoding/json"
	"net/http"

	customerrorconst "github.com/dcssoftware/bafoeg-manager/src/helper/debug/customerrors/custom-error-const"
	"github.com/dcssoftware/bafoeg-manager/src/static/assets"
	"github.com/gofiber/fiber/v3"
)

// @Summary		Get the dependent software licenses - json
// @Description	Retrieves the software licenses
// @Tags			assets
//
// @Produce		application/json
//
// @Success		200	{file}		file							"A GIF image of a gopher drinking coffee."
// @Failure		404	{object}	swagger.NotFoundErrorResponse	"code":	404,	"message":	"Not found ❌"
//
// @Router			/api/asset/licenses [get]
func (h *GeneralHandler) AssetLicenses(c fiber.Ctx) error {
	isValid := json.Valid(assets.SoftwareLicenses)
	if !isValid {
		return c.Status(http.StatusInternalServerError).SendString(customerrorconst.INTERNAL_SERVER_ERROR_MESSAGE)
	}
	c.Set("Content-Type", "application/json")
	return c.Send(assets.SoftwareLicenses)
}
