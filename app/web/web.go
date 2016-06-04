package web

import (
	"github.com/lcaballero/reverb-example/cli"
	"github.com/lcaballero/hitman"
	"github.com/lcaballero/reverb/base"
	"github.com/labstack/echo/engine/standard"
	"github.com/lcaballero/reverb-example/app/view"
	"github.com/lcaballero/gel"
	"log"
	app "github.com/lcaballero/reverb-example/app/context"
)

type WebServer struct{
	*cli.Args
	IncludesResolver func(string) gel.View
}

func NewWebServer(args *cli.Args) (*WebServer, error) {
	ws := &WebServer{
		Args: args,
	}
	return ws, nil
}

func (w *WebServer) Start() hitman.KillChannel {
	done := hitman.NewKillChannel()
	go func() {
		go w.run()
		for {
			select {
			case cleaner := <-done:
				cleaner.WaitGroup.Done()
				return
			}
		}
	}()
	return done
}

func (w *WebServer) run() {
	log.Printf("finding templates at: %s", w.IncludesDir)
	log.Printf("using assets found at: %s", w.AssetsDir)

	ctx := app.NewContext(w.Args)

	r := base.NewRegister()
	r.Get("/", view.Index(ctx))
	r.Get("/asset/:kind/:hash/:file", ctx.ToHandler())

	log.Printf("starting web server at: %s", w.Ip)
	r.Echo.Run(standard.New(w.Ip))
}
