package stdext

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ForEachLine reads lines from r and applies f to each of them. The string
// arguments to f will not include line-endings.
func ForEachLine(r io.Reader, f func(string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	return scanner.Err()
}

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
