package main

import (
	"github.com/ronelliott/go-cli"

	"github.com/ronelliott/go-events/commands/events/subcommands"
)

func main() {
	cli.RunWithErrorsAndExit(&cli.AppConfig{
		Name: "events",
		Desc: "Events service and utilities",
		Help: "Events service and utilities",
		Commands: []cli.CommandConverter{
			&cli.CommandConfig{
				Name: "client",
				Desc: "Run a client instance",
				Help: "Run a client instance",
				Run:  &subcommands.ClientCmd{},
			},
			&cli.CommandConfig{
				Name: "server",
				Desc: "Run a server instance",
				Help: "Run a server instance",
				Run:  &subcommands.ServerCmd{},
			},
		},
	})
}
