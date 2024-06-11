//go:build darwin

package gotracing

import (
	"os"
	"strings"
	"syscall"
)

func supportsColor() bool {
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, os.Stdout.Fd(), uintptr(syscall.TIOCGWINSZ), 0); err == 0 {
		term := strings.ToLower(os.Getenv("TERM"))
		if strings.Contains(term, "color") || strings.Contains(term, "xterm") || strings.Contains(term, "screen") || strings.Contains(term, "vt100") {
			return true
		}
	}
	return false
}
