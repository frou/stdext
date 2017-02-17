package stdext

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	OwnerWritableDir = 0755
	OwnerWritableReg = 0644
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

// ExecutableBasename returns the basename of the current executable (i.e.
// argv0 without its full path).
//
// TODO(DH): Go 1.8 introduced the following:
// https://golang.org/pkg/os/#Executable
func ExecutableBasename() string {
	return filepath.Base(os.Args[0])
}
