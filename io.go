package stdext

import (
	"io"
)

// Close x, assigning the error return value to the blank identifier. This is
// for keeping static analysis tools like `errcheck` happy in situations where
// you conclude that an error return value can safely be ignored. Particularly
// useful with defer, e.g. defer stdext.Close(f)
func Close(x io.Closer) {
	_ = x.Close()
}
