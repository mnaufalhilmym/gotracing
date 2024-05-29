package gotracing

import (
	"runtime"
)

type stacktrace struct {
	kind string
	msg  []any
	file string
	fn   string
	line int
}

type Stacktraces []stacktrace

func newStacktraces(level Level, msg ...any) Stacktraces {
	kind := ""
	switch level {
	case 0: // trace
		kind += termColorPurple + "TRACE" + termColorReset
	case 1: // debug
		kind += termColorBlue + "DEBUG" + termColorReset
	case 2: // info
		kind += termColorGreen + "INFO" + termColorReset
	case 3: // warn
		kind += termColorYellow + "WARN" + termColorReset
	case 4: // error
		kind += termColorRed + "ERROR" + termColorReset
	}
	return Stacktraces{{kind: kind, msg: msg}}
}

func traceStack(maxPC uint, stacks Stacktraces) Stacktraces {
	pcs := make([]uintptr, maxPC)
	pcEntries := runtime.Callers(3, pcs)
	if pcEntries == 0 {
		return stacks
	}

	frames := runtime.CallersFrames(pcs)

	for {
		frame, more := frames.Next()

		stacks[len(stacks)-1].file = frame.File
		stacks[len(stacks)-1].fn = frame.Func.Name()
		stacks[len(stacks)-1].line = frame.Line

		if more {
			stacks = append(stacks, stacktrace{})
		} else {
			break
		}
	}

	return stacks
}
