package dev

import (
	"sheep"
	"sheep/internal/docker"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func newStopCommand(app *sheep.App) *cli.Command {
	return &cli.Command{
		Name:  "stop",
		Usage: "stop a Mooncard's dependency container",
		Action: func(ctx *cli.Context) error {
			name := ctx.Args().Get(0)
			// TODO check if dep exists

			c, err := docker.NewClient()
			if err != nil {
				return err
			}

			logger := logrus.WithField("container", name)
			logger.Info("looking for container")
			container, err := c.FindContainer(ctx.Context, name)
			if err != nil {
				logger.Error("unable to find the container")
				return err
			}

			if container.State == "exited" {
				logger.Info("dependency already down")
				return nil
			}

			logger.Info("stopping container")
			if err := c.ContainerStop(ctx.Context, container.ID, nil); err != nil {
				logger.Info("unable to stop container")
				return nil
			}
			logger.Info("container stopped")

			return nil
		},
	}
}
