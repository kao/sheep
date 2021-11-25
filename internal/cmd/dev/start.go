package dev

import (
	"context"
	"errors"
	"sheep"
	"sheep/internal/docker"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func newStartCommand(app *sheep.App) *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "start a Mooncard's dependency container",
		Action: func(ctx *cli.Context) error {
			name := ctx.Args().Get(0)
			dep := app.DependenciesCfg[name]
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
				if container.Status != "" && strings.Contains(container.Status, "Up") {
					logger.Error("container already running")
					return errors.New("container already running")
				}
				containerID = container.ID
			}

			logger.Info("starting container")

			if err := c.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{}); err != nil {
				logger.Info("unable to start container")
				return nil
			}

			logger.Info("container started")

			printDependencyInformation(dep)

			return nil
		},
	}
}
