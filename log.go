package stdext

import (
	"io"
	"log"
	"os"
)

// MirrorLogToFile arranges for the standard logger to write to a file *in
// addition* to the default destination (stderr). The standard logger is the
// one used by calling e.g. log.Print(...)
func MirrorLogToFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		OwnerWritableReg)
	if err != nil {
		return nil, err
	}
	log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	return logFile, nil
}
