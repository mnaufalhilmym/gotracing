//go:build linux

package gotracing

import (
	"os"
	"strings"
	"syscall"
	"unsafe"
)

func supportsColor() bool {
	var termios syscall.Termios
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, os.Stdout.Fd(), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&termios)), 0, 0, 0); err == 0 {
		term := strings.ToLower(os.Getenv("TERM"))
		if strings.Contains(term, "color") || strings.Contains(term, "xterm") || strings.Contains(term, "screen") || strings.Contains(term, "vt100") {
			return true
		}
	}
	return false
}
