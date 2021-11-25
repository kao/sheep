package main

import (
	"log"
	"os"
	"sheep"
	"sheep/internal/cmd"
)

func main() {
	app := sheep.NewApp()
	app.Commands = cmd.NewCommands(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
