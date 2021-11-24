package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"mcc"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type dockerContainer struct {
	name        string
	image       string
	ports       nat.PortMap
	env         []string
	information []string
}

var dependenciesMap = map[string]*dockerContainer{
	"postgres": {
		name:  "mooncard-postgres",
		image: "postgres:12-alpine",
		ports: nat.PortMap{
			"5432/tcp": {
				{HostPort: "5432"},
			},
		},
		env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=password",
		},

		information: []string{
			"URI: postgres:password@localhost:5432/postgres",
		},
	},
}

func main() {
	app := cli.NewApp()
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "version",
		Usage: "display mcc version",
		Action: func(c *cli.Context) error {
			fmt.Println(mcc.Version)
			return nil
		},
	})

	app.Commands = append(app.Commands, &cli.Command{
		Name:    "dev",
		Aliases: []string{},
		Usage:   "start Mooncard's dependencies containers",
		Action: func(ctx *cli.Context) error {
			c, err := client.NewClientWithOpts(client.FromEnv)
			if err != nil {
				return err
			}

			name := "postgres"
			dep := dependenciesMap[name]
			logger := logrus.WithField("container", name)
			logger.Info("looking for container")

			containers, err := c.ContainerList(
				context.Background(),
				types.ContainerListOptions{
					Limit:   1,
					Filters: filters.NewArgs(filters.Arg("name", name)),
				},
			)
			if err != nil {
				logger.Warn("unable to list containers")
				return err
			}

			var containerID string
			if len(containers) == 0 {
				logger.Info("container not found")
				containerID = ""
			} else {
				containerID = containers[0].ID
			}

			if containerID == "" {
				logger.Info("pulling image")
				out, err := c.ImagePull(context.Background(), dep.image, types.ImagePullOptions{})
				if err != nil {
					logger.Error("unable to pull Docker image %s")
					return nil
				}
				io.Copy(os.Stdout, out)

				logger.Info("creating Docker container")
				cont, err := c.ContainerCreate(
					context.Background(),
					&container.Config{
						Image: dep.image,
						Env:   dep.env,
					},
					&container.HostConfig{PortBindings: dep.ports},
					nil,
					nil,
					name)
				if err != nil {
					logger.Error("unable to create container")
					logger.Error(err)
					return nil
				}

				containerID = cont.ID
			}

			logger.Info("starting Docker image")
			if err := c.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{}); err != nil {
				logger.Info("unable to start container")
				return nil
			}

			return nil
		},
	})

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
