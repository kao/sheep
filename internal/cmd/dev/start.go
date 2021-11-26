package dev

import (
	"sheep"
	"sheep/internal/docker"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func newStartCommand(app *sheep.App) *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "start a Sheep's dependency",
		Action: func(ctx *cli.Context) error {
			c, err := docker.NewClient()
			if err != nil {
				return err
			}

			name := ctx.Args().Get(0)
			for _, d := range app.DependenciesCfg {
				if name != "" && name != strings.TrimPrefix(d.Name, "sheep-") {
					continue
				}

				logger := logrus.WithField("dependency", d.Name)
				logger.Info("starting dependency")

				if err := c.StartContainer(ctx.Context, d); err != nil {
					err = errors.Wrap(err, "unable to start dependency")

					if name != "" && name != strings.TrimPrefix(d.Name, "sheep-") {
						return err
					} else {
						logger.Error(err)
						continue
					}
				}

				logger.Info("dependency started")

				printDependencyInformation(d)
			}

			return nil
		},
	}
}
