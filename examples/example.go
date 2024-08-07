package main

import (
	"os"

	"github.com/x1um1n/askme"
)

func main() {
	if askme.AskYN("Is the sky blue?") {
		println("Correct!")
	} else {
		switch askme.Ask("Are you feeling ok?") {
		case askme.Yes:
			println("I'm glad to hear that!")
		case askme.No:
			println("I sense a deep dejection in your diodes Robot. It saddens me & I globber.")
		case askme.Skip:
			os.Exit(1)
		case askme.Quit:
			os.Exit(0)
		}
	}
}
