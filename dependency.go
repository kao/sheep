package sheep

import "github.com/docker/go-connections/nat"

type Dependency struct {
	Name        string
	Image       string
	Ports       nat.PortMap
	Env         []string
	Information []string
}

var dependenciesMap = map[string]*Dependency{
	"postgres": {
		Name:  "mooncard-postgres",
		Image: "postgres:12-alpine",
		Ports: nat.PortMap{
			"5432/tcp": {
				{HostPort: "5432"},
			},
		},
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=password",
		},
		Information: []string{
			"URI: postgres:password@localhost:5432/postgres",
		},
	},
	"redis": {
		Name:  "mooncard-redis",
		Image: "redis:6.2-alpine",
		Ports: nat.PortMap{
			"6379/tcp": {
				{HostPort: "6379"},
			},
		},
		Env: []string{},
		Information: []string{
			"Listener: localhost:6379",
		},
	},
}
