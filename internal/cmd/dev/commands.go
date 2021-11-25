package dev

import "github.com/urfave/cli/v2"

func NewDevCommands() *cli.Command {
	return &cli.Command{
		Name:  "dev",
		Usage: "Commands to interact with dev environment",
		Subcommands: []*cli.Command{
			newUpCommand(),
			newDownCommand(),
		},
	}
}
