package dev

import (
	"fmt"
	"sheep"
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func NewInformationCommand(app *sheep.App) *cli.Command {
	return &cli.Command{
		Name:  "info",
		Usage: "show Sheep's dependency information",
		Action: func(ctx *cli.Context) error {
			name := ctx.Args().Get(0)

			if name == "" {
				for _, d := range app.DependenciesCfg {
					printDependencyInformation(d)
				}
			} else {
				dep := app.DependenciesCfg[name]
				if dep == nil {
					return errors.New("unknown dependency")
				}

				printDependencyInformation(dep)
			}

			return nil
		},
	}
}

func printDependencyInformation(d *sheep.Dependency) {
	if len(d.Information) == 0 {
		return
	}

	fmt.Printf("%s\n", strings.TrimPrefix(d.Name, "sheep-"))
	for _, i := range d.Information {
		fmt.Printf("  -> %s\n", i)
	}
}
