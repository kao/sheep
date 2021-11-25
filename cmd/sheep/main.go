package main

import (
	"log"
	"os"
	"sheep/internal/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: cmd.NewCommands(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
