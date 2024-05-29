package gotracing

import (
	"os"
	"syscall"
	"unsafe"
)

func supportsColor() bool {
	var termios syscall.Termios
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, os.Stdout.Fd(), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&termios)), 0, 0, 0); err == 0 {
		term := os.Getenv("TERM")
		if term == "xterm-256color" || term == "screen-256color" || term == "tmux-256color" || term == "rxvt-unicode-256color" {
			return true
		}
		if os.Getenv("COLORTERM") == "truecolor" {
			return true
		}
	}
	return false
}
