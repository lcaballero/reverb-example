package view

import (
	"github.com/lcaballero/reverb/view"
	"github.com/labstack/echo"
	"net/http"
	. "github.com/lcaballero/gel"
	app "github.com/lcaballero/reverb-example/app/context"
)

func Index(ctx *app.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := view.Root{
			Assets: Frag(
				ctx.ToAsset("reset.css"),
			),
			Scripts: Frag(
				ctx.ToAsset("hello-world.js"),
			),
			Main: Div.Class("theme1")(
				ctx.Include("hello-world.html"),
			),
		}
		return ctx.ToHtml(c, http.StatusOK, r)
	}
}
