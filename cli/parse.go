package cli

import (
	"github.com/jessevdk/go-flags"
	"os"
)

type Args struct {
	Ip string `long:"ip" description:"The ip:host which the server should bind to." default:"127.0.0.1:4000"`
	AssetsDir string `short:"a" long:"assets" description:"Directory from which to server static assets."`
	IncludesDir string `long:"includes" description:"Directory from which include templates."`
}

func ParseArgs(data interface{}) *Args {
	args := &Args{}
	parser := flags.NewParser(args, flags.Default)
	_, err := parser.ParseArgs(os.Args)
	if err != nil {
		os.Exit(1)
	}
	return args
}
