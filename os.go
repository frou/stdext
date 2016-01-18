package stdext

import (
	"fmt"
	"os"
)

// Exit exits the current process. If failure is non-nil, it is printed to
// standard error and the exit status is non-zero.
func Exit(failure interface{}) {
	var status int
	if failure != nil {
		fmt.Fprintln(os.Stderr, failure)
		status = 3
	}
	os.Exit(status)
}
