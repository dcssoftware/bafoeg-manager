package http

import (
	"github.com/dcssoftware/bafoeg-manager/src/static/assets"
	"github.com/gofiber/fiber/v3"
)

// @Summary		Get the website logo
// @Description	Retrieves the website logo
// @Tags			assets
//
// @Produce		image/png
// @Success		200	{file}		file							"Website Logo."
// @Failure		404	{object}	swagger.NotFoundErrorResponse	"code":	404,	"message":	"Not found ❌"
//
// @Router			/api/asset/logo-branding [get]
func (h *GeneralHandler) AssetLogoBranding(c fiber.Ctx) error {
	c.Set("Content-Type", "image/png")
	return c.SendString(assets.LogoBranding)
}
