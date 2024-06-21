package gotracing

import (
	"runtime"
)

type Traces struct {
	level       Level
	msg         []any
	stacktraces []stacktrace
}

type stacktrace struct {
	file string
	fn   string
	line int
}

func traceStack(maxPC uint, traces Traces) Traces {
	pcs := make([]uintptr, maxPC)
	pcEntries := runtime.Callers(3, pcs)
	if pcEntries == 0 {
		return traces
	}

	stacks := make([]stacktrace, 0, maxPC)

	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()

		stacks = append(stacks, stacktrace{
			file: frame.File,
			fn:   frame.Func.Name(),
			line: frame.Line,
		})

		if !more {
			break
		}
	}

	traces.stacktraces = stacks

	return traces
}
