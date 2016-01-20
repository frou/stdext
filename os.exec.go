package stdext

import (
	"os/exec"
	"runtime"
)

// Launch attempts to open noun as if it were launched in the graphical
// environment of the host operating system.
//
// Some examples:
//
//   Launch("path/to/file.pdf")
//   Likely opens that file in Preview.app on OS X
//
//   Launch("x.txt")
//   Likely opens that file in Notepad.exe on Windows
//
//   Launch("http://repl.it/")
//   Likely opens that URL in the default web browser on Linux, Windows & OS X
func Launch(noun string) error {
	return exec.Command(osLauncher(), noun).Run()
}

func osLauncher() string {
	switch runtime.GOOS {
	case "darwin":
		return "open"
	case "windows":
		return "start"
	default:
		return "xdg-open"
	}
}
