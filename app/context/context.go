package context

import (
	"github.com/lcaballero/reverb/base"
	"github.com/lcaballero/gel"
	"github.com/lcaballero/reverb/view"
	"github.com/labstack/echo"
	"path/filepath"
	"github.com/lcaballero/reverb-example/cli"
	"io/ioutil"
)

type Context struct {
	ToAsset func(filename string) gel.View
	ToHandler func() echo.HandlerFunc
	ToHtml func(c echo.Context, httpCode int, view gel.Viewable) error
	Include func(file string) gel.View
}


func NewContext(args *cli.Args) *Context {
	includes := func(file string) string {
		return filepath.Join(args.IncludesDir, file)
	}

	resolver := view.RootResolver(args.AssetsDir)
	html := base.RenderHtml{}

	ctx := &Context{
		ToHtml: html.ToHtml,
		ToAsset: resolver.ToAsset,
		ToHandler: resolver.ToHandler,
		Include: gel.NewInserter(includes, ioutil.ReadFile).Include,
	}
	return ctx
}