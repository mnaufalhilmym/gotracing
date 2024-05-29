package gotracing

import (
	"runtime"
)

type stacktrace struct {
	level Level
	msg   []any
	file  string
	fn    string
	line  int
}

type Stacktraces []stacktrace

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
