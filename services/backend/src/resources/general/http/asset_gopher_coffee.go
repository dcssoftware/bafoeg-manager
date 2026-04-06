package http

import (
	"github.com/dcssoftware/bafoeg-manager/src/static/assets"
	"github.com/gofiber/fiber/v3"
)

// @Summary		Get a nice gopher trinking coffee - image
// @Description	Retrieves a GIF image of a gopher enjoying coffee.
// @Tags			assets
//
// @Produce		image/gif
//
// @Success		200	{file}		file							"A GIF image of a gopher drinking coffee."
// @Failure		404	{object}	swagger.NotFoundErrorResponse	"code":	404,	"message":	"Not found ❌"
//
// @Router			/api/asset/gopher-coffee [get]
func (h *GeneralHandler) AssetGopherCoffee(c fiber.Ctx) error {
	c.Set("Content-Type", "image/gif")
	return c.SendString(assets.GopherCoffee)
}
