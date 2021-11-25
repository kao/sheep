package dev

import (
	"context"
	"mcc/internal/docker"

	"github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func newUpCommand() *cli.Command {
	return &cli.Command{
		Name:  "up",
		Usage: "start a Mooncard's dependency container",
		Action: func(ctx *cli.Context) error {
			name := ctx.Args().Get(0)
			dep := docker.DependenciesMap[name]
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

			var containerID string
			if container == nil {
				logger.Info("pulling image and creating the container")
				containerID, err = c.PullImageAndCreateContainer(ctx.Context, dep)
				if err != nil {
					return err
				}
			} else {
				containerID = container.ID
			}

			logger.Info("starting container")
			if err := c.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{}); err != nil {
				logger.Info("unable to start container")
				return nil
			}
			logger.Info("container started")

			return nil
		},
	}
}
