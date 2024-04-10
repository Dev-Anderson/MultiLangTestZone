package newlogger

import (
	"os"
	"runtime"
)

func isRunningInDockerContainer() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}

func isRunningInLinux() bool {
	return runtime.GOOS == "linux"
}
