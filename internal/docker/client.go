package docker

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

type Client struct {
	*client.Client
}

func NewClient() (*Client, error) {
	c, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &Client{c}, nil
}

func (c *Client) FindContainer(ctx context.Context, name string) (*types.Container, error) {
	containers, err := c.ContainerList(
		ctx,
		types.ContainerListOptions{
			Limit:   1,
			Filters: filters.NewArgs(filters.Arg("name", name)),
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list containers")
	}

	if len(containers) == 0 {
		return nil, nil
	}

	return &containers[0], nil
}

func (c *Client) PullImageAndCreateContainer(ctx context.Context, dep *dockerContainer) (string, error) {
	out, err := c.ImagePull(ctx, dep.image, types.ImagePullOptions{})
	if err != nil {
		return "", errors.Wrap(err, "unable to pull Docker image")
	}
	io.Copy(os.Stdout, out)

	cont, err := c.ContainerCreate(
		ctx,
		&container.Config{
			Image: dep.image,
			Env:   dep.env,
		},
		&container.HostConfig{PortBindings: dep.ports},
		nil,
		nil,
		dep.name)
	if err != nil {
		return "", errors.Wrap(err, "unable to create container")
	}

	return cont.ID, nil
}
