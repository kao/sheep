package main

import (
	"fmt"
	"log"
	"mcc"
	"mcc/internal/cmd/dev"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "display mcc version",
				Action: func(c *cli.Context) error {
					fmt.Println(mcc.Version)
					return nil
				},
			},
			dev.NewDevCommands(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
