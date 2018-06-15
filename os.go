package stdext

import (
	"fmt"
	"os"
	"path/filepath"
)

// Octal values to use for e.g. the perm parameter of os.OpenFile.
const (
	OwnerWritableDir = 0755
	OwnerWritableReg = 0644
)

// Exit exits the current process. If failure is nil, the process exit code
// will be zero. If failure is an int, the process exit code will be that
// value. Otherwise, the value of failure will be printed to stderr and the
// process exit code will be non-zero.
func Exit(failure interface{}) {
	switch failure := failure.(type) {
	case nil:
		os.Exit(0)
	case int:
		os.Exit(failure)
	case error:
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", failure)
		os.Exit(2)
	default:
		fmt.Fprintln(os.Stderr, failure)
		os.Exit(3)
	}
}

// ExecutableBasename returns the basename of the current executable (i.e.
// argv0 without its full path).
func ExecutableBasename() string {
	// We could use os.Executable() instead of os.Args[0] , but it can fail.
	return filepath.Base(os.Args[0])
}
