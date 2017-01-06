package stdext

import (
	"bufio"
	"fmt"
	"os"
)

// PressReturnTo prints a message with the given verb to stderr, then returns
// after reading a line from stdin (which is discarded).
func PressReturnTo(verb string) {
	fmt.Fprintln(os.Stderr, "Press return to", verb)
	bufio.NewScanner(os.Stdin).Scan()
}

// PressReturnToContinue prints a message to stderr, then returns after reading
// a line from stdin (which is discarded).
func PressReturnToContinue() {
	PressReturnTo("continue")
}
