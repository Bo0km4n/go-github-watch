package main

import (
	"os"

	"github.com/Bo0km4n/go-github-watch"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	ggw.ParseArgs()

	app := cli.NewApp()
	app.Name = ggw.Name
	app.Usage = ggw.Usage
	app.Version = ggw.Version
	app.Commands = initCommands()
	app.Run(os.Args)
}
