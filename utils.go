package ggw

import (
	"os"
)

const (
	// Name is the program name
	Name = "GGW"
	// Usage is for simple description
	Usage = "Manage your Github watching list"
)

func ParseArgs() {
	if len(os.Args) == 1 {
		displayLogo()
	} else if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "h" || os.Args[1] == "help" {
			displayLogo()
		}
	}
}
