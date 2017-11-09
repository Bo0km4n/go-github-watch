package main

import (
	"fmt"

	cli "gopkg.in/urfave/cli.v1"
)

func initialize(c *cli.Context) error {
	fmt.Println("Initilize command")
	return nil
}

func list(c *cli.Context) error {
	fmt.Println("ls list")
	return nil
}
