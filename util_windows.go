//go:build windows

package gotracing

import (
	"os"

	"golang.org/x/sys/windows"
)

const (
	ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004
	STD_OUTPUT_HANDLE                  = uint32(windows.STD_OUTPUT_HANDLE)
)

func supportsColor() bool {
	handle := windows.Handle(os.Stdout.Fd())
	var mode uint32
	if err := windows.GetConsoleMode(handle, &mode); err == nil {
		mode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING
		if err := windows.SetConsoleMode(handle, mode); err == nil {
			return true
		}
	}
	return false
}
