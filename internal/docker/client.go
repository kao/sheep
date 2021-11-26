package docker

import (
	"context"
	"io"
	"os"
	"sheep"
	"strings"

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

func (c *Client) StartContainer(ctx context.Context, dep *sheep.Dependency) error {
	container, err := c.FindContainer(ctx, dep.Name)
	if err != nil {
		return errors.Wrap(err, "unable to find the container")
	}

	var containerID string
	if container == nil {
		containerID, err = c.PullImageAndCreateContainer(ctx, dep)
		if err != nil {
			return err
		}
	} else {
		if container.Status != "" && strings.Contains(container.Status, "Up") {
			return errors.New("container already running")
		}
		containerID = container.ID
	}

	if err := c.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		return errors.Wrap(err, "unable to start container")
	}

	return nil
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

func (c *Client) PullImageAndCreateContainer(ctx context.Context, dep *sheep.Dependency) (string, error) {
	out, err := c.ImagePull(ctx, dep.Image, types.ImagePullOptions{})
	if err != nil {
		return "", errors.Wrap(err, "unable to pull Docker image")
	}
	io.Copy(os.Stdout, out)

	cont, err := c.ContainerCreate(
		ctx,
		&container.Config{
			Image: dep.Image,
			Env:   dep.Env,
		},
		&container.HostConfig{PortBindings: dep.Ports},
		nil,
		nil,
		dep.Name)
	if err != nil {
		return "", errors.Wrap(err, "unable to create container")
	}

	return cont.ID, nil
}

func (c *Client) StopContainer(ctx context.Context, name string) error {
	container, err := c.FindContainer(ctx, name)
	if err != nil {
		return errors.Wrap(err, "unable to find the container")
	}

	if container.State == "exited" {
		return errors.New("dependency already stopped")
	}

	if err := c.ContainerStop(ctx, container.ID, nil); err != nil {
		return errors.Wrap(err, "unable to stop container")
	}

	return nil
}
