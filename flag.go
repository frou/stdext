package stdext

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// SetPreFlagsUsageMessage replaces the standard usage message used by package
// flag with one that, in addition to describing the flags that are defined,
// also has a general textual description of the command, followed by 0 or more
// examples of the non-flag argument(s) expected by the command.
//
// The examples will be printed prefixed by "CMD [flags] " where CMD is the
// value of os.Args[0].
//
// This function should be called before flag.Parse.
func SetPreFlagsUsageMessage(desc string, examples ...string) {
	flag.Usage = func() {
		if desc != "" {
			desc = strings.TrimSpace(desc)
			// Un-hardwrap desc but maintain paragraph breaks.
			desc = strings.Replace(desc, "\n\n", "\035", -1)
			desc = strings.Replace(desc, "\n", " ", -1)
			desc = strings.Replace(desc, "\035", "\n\n", -1)
			fmt.Fprintln(os.Stderr, desc)
			fmt.Fprintln(os.Stderr)
		}
		fmt.Fprintln(os.Stderr, "usage:")
		prefix := "  " + os.Args[0] + " [flags]"
		if len(examples) == 0 {
			fmt.Fprintln(os.Stderr, prefix)
		} else {
			for _, e := range examples {
				fmt.Fprintln(os.Stderr, prefix, e)
			}
		}
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "flags:")
		flag.PrintDefaults()
	}
}

// ParseFlagsExpectingNArgs calls flag.Parse and then checks that n non-flag
// arguments are present in addition to the parsed flags. If not, flag.Usage is
// called and an error is returned.
func ParseFlagsExpectingNArgs(n int) error {
	flag.Parse()
	if got := flag.NArg(); got != n {
		flag.Usage()
		fmt.Fprintln(os.Stderr)
		return fmt.Errorf(
			"Expected %d non-flag argument(s) but got %d", n, got)
	}
	return nil
}
