package start

import (
	"github.com/lcaballero/reverb-example/cli"
	"github.com/lcaballero/reverb-example/app/web"
	"github.com/lcaballero/hitman"
	"github.com/vrecan/death"
	"syscall"
)

func Start() {
	args := cli.ParseArgs(&cli.Args{})

	targets := hitman.NewTargets()
	targets.AddOrPanic(web.NewWebServer(args))

	death.NewDeath(syscall.SIGTERM, syscall.SIGINT).WaitForDeath(targets)
}

