package main

import cli "gopkg.in/urfave/cli.v1"

func initCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialize your GitHub Status",
			Action:  initialize,
		},
		{
			Name:    "ls",
			Aliases: []string{"l"},
			Usage:   "Show your watching repos",
			Action:  list,
		},
	}
}
