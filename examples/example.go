package main

import (
	"fmt"
	"os"

	"github.com/x1um1n/askme"
)

func main() {
	// Yes/No example usage
	if askme.AskYN("Is the sky blue?") {
		fmt.Println("Correct!")
	} else {
		// Yes/No/Quit/Skip example usage
		switch askme.Ask("Are you feeling ok?") {
		case askme.Yes:
			fmt.Println("I'm glad to hear that!")
		case askme.No:
			fmt.Println("I sense a deep dejection in your diodes Robot. It saddens me & I globber.")
		case askme.Skip:
			os.Exit(1)
		case askme.Quit:
			os.Exit(0)
		}
	}
}
