package sheep

import "github.com/urfave/cli/v2"

const Version = "0.0.1"

type App struct {
	*cli.App
	DependenciesCfg map[string]*Dependency
}

func NewApp() *App {
	cliApp := cli.NewApp()
	return &App{
		cliApp,
		dependenciesMap,
	}
}
