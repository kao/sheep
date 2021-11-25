package docker

import "github.com/docker/go-connections/nat"

type dockerContainer struct {
	name        string
	image       string
	ports       nat.PortMap
	env         []string
	information []string
}

var DependenciesMap = map[string]*dockerContainer{
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
