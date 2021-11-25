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
	"rabbitmq": {
		Name:  "mooncard-rabbitmq",
		Image: "rabbitmq:3.9.10-management-alpine",
		Ports: nat.PortMap{
			"5672/tcp": {
				{HostPort: "5672"},
			},
			"15672/tcp": {
				{HostPort: "15672"},
			},
		},
		Information: []string{
			"Listener: localhost:5672",
			"WebUI: http://localhost:15672",
			"Credentials: guest/guest",
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
