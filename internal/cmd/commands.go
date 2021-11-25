package cmd

import (
	"fmt"
	"sheep"
	"sheep/internal/cmd/dev"

	"github.com/urfave/cli/v2"
)

func NewCommands(app *sheep.App) []*cli.Command {
	return []*cli.Command{
		newVersionCommand(),
		dev.NewDevCommands(app),
	}
}

func newVersionCommand() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "display sheep version",
		Action: func(c *cli.Context) error {
			ascii :=
				`
		      __                __________
	    ,''''--'''  ''-''-.        (          )
	  ,'            ,-- ,-'.      (            )
	 (//            '"'| 'a \    (     %s    )
	   |    ';         |--._/     (            )
	   \    _;-._,    /            (__________)
	    \__/\\   \__,'
	     ||  ''   \|\\
	     \\        \\''
	      ''        ''
`
			fmt.Printf(ascii, sheep.Version)
			return nil
		},
	}
}
