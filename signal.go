package stdext

import (
	"os"
	"os/signal"
)

// HandleSignal sets a handler for sig and specifies whether it should handle
// the signal repeatedly as opposed to only once.
func HandleSignal(sig os.Signal, repeatedly bool, handler func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, sig)
	go func() {
		defer signal.Stop(c)
		for {
			<-c
			handler()
			if !repeatedly {
				break
			}
		}
	}()
}

// SignalSelf sends sig to the current process.
func SignalSelf(sig os.Signal) error {
	self, err := os.FindProcess(os.Getpid())
	if err != nil {
		return err
	}
	return self.Signal(sig)
}
