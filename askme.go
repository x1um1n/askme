package askme

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/x1um1n/checkerr"
)

type Answer int32

const (
	Yes Answer = iota
	No
	Skip
	Quit
)

var (
	quitResponses []string = []string{"q", "quit", "exit"}
	skipResponses []string = []string{"s", "skip", "skipit"}
	noResponses   []string = []string{"n", "no", "nope", "nah", "nay"}
	yesResponses  []string = []string{"y", "yes", "yup", "yep", "yessirreebob", "ya", "yarp", "yarr", "ja", "aye", "ochaye", "okthen"}
)

// Ask prints a string and asks the user for confirmation, giving options for "yes", "no", "skip", and "quit".
// If multiple responses are submitted, it will match against "quit", "skip" and "no" ahead of "yes"
func Ask(s string) Answer {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [yes/no/skip/quit]: ", s)

		response, err := reader.ReadString('\n')
		checkerr.CheckFatal(err) // script exits if it fails to read the response

		response = strings.ToLower(strings.TrimSpace(response))

		if slices.Contains(quitResponses, response) {
			return Quit
		} else if slices.Contains(skipResponses, response) {
			return Skip
		} else if slices.Contains(noResponses, response) {
			return No
		} else if slices.Contains(yesResponses, response) {
			return Yes
		}
	}
}

// askForConfirmation asks the user for confirmation. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y" and "n" will work too.
// It returns true if the user typed "yes", false otherwise.
func AskYN(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [yes/no]: ", s)

		response, err := reader.ReadString('\n')
		checkerr.CheckFatal(err)

		response = strings.ToLower(strings.TrimSpace(response))

		if slices.Contains(noResponses, response) {
			return false
		} else if slices.Contains(yesResponses, response) {
			return true
		}
	}
}
