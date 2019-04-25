package stdext

import (
	"fmt"
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
	launcher, _ := osLauncher()
	return exec.Command(launcher, noun).Run()
}

func LaunchInBackground(noun string) error {
	launcher, backgroundFlag := osLauncher()
	if backgroundFlag == "" {
		return fmt.Errorf("background flag for '%v' command is unknown", launcher)
	}
	return exec.Command(launcher, backgroundFlag, noun).Run()
}

func osLauncher() (command string, backgroundFlag string) {
	switch runtime.GOOS {
	case "darwin":
		return "open", "-g"
	case "windows":
		return "start", ""
	default:
		return "xdg-open", ""
	}
}
