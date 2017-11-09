package ggw

import "github.com/fatih/color"

var (
	// Version is the default version of SKM
	Version = "0.1"
	logo    = `
 ________  ___       __   ___       __      
|\   ____\|\  \     |\  \|\  \     |\  \    
\ \  \___|\ \  \    \ \  \ \  \    \ \  \   
 \ \  \  __\ \  \  __\ \  \ \  \  __\ \  \  
  \ \  \|\  \ \  \|\__\_\  \ \  \|\__\_\  \ 
   \ \_______\ \____________\ \____________\
    \|_______|\|____________|\|____________|
GWW V%s
https://github.com/Bo0km4n/go-github-watch

`
)

func displayLogo() {
	color.Cyan(logo, Version)
}
