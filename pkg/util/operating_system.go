package util

import (
	"fmt"
	"os/exec"
	"runtime"
)

// opens the URL in the native browser
func OpenInBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "linux":
		cmd = "xdg-open"
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	default:
		return fmt.Errorf("unsupported platform")
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
