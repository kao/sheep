package dev

import (
	"sheep"
	"sheep/internal/docker"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func newStopCommand(app *sheep.App) *cli.Command {
	return &cli.Command{
		Name:  "stop",
		Usage: "stop a Sheep's dependency",
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

				logger.Info("stopping dependency")
				if err := c.StopContainer(ctx.Context, d.Name); err != nil {
					err = errors.Wrap(err, "unable to stop the dependency")

					if name != "" && name != strings.TrimPrefix(d.Name, "sheep-") {
						return err
					} else {
						logger.Error(err)
						continue
					}
				}
				logger.Info("dependency stopped")
			}

			return nil
		},
	}
}
