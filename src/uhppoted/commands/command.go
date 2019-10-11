package commands

import (
	"flag"
	"os"
)

type Context struct {
}

type Command interface {
	Execute(context Context) error
	Cmd() string
	Description() string
	Usage() string
	Help()
}

var VERSION = "v0.04.0"

var cli = []Command{
	&Daemonize{},
	&Undaemonize{},
	&Version{VERSION},
	&Help{},
}

func Parse() (Command, error) {
	var cmd Command = nil
	var err error = nil

	if len(os.Args) > 1 {
		for _, c := range cli {
			if c.Cmd() == flag.Arg(0) {
				cmd = c
			}
		}
	}

	return cmd, err
}