package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Version = "0.0.1"

var BuildFlag string

var VersionCmd = &cli.Command{
	Name:    "version",
	Usage:   "print version",
	Aliases: []string{"V"},
	Action: func(_ *cli.Context) error {
		fmt.Println(Version + "+" + BuildFlag)
		return nil
	},
}
