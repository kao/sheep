package dev

import (
	"sheep"

	"github.com/urfave/cli/v2"
)

func NewDevCommands(app *sheep.App) *cli.Command {
	return &cli.Command{
		Name:  "dev",
		Usage: "Commands to interact with dev environment",
		Subcommands: []*cli.Command{
			newStartCommand(app),
			newStopCommand(app),
		},
	}
}
