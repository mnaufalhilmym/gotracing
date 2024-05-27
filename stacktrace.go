package gotracing

import (
	"fmt"
	"runtime"
	"time"
)

type stacktrace struct {
	msg  []any
	file string
	fn   string
	line int
}

func traceStack(maxPC uint, msg ...any) []stacktrace {
	pcs := make([]uintptr, maxPC)
	pcEntries := runtime.Callers(3, pcs)
	if pcEntries == 0 {
		return nil
	}

	frames := runtime.CallersFrames(pcs)

	stacks := []stacktrace{{msg: msg}}

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

func printStack(l level, st []stacktrace) {
	kind := ""
	switch l {
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
	for i, s := range st {
		if s.msg != nil {
			fmt.Printf("%s  %s: ", termColorComment+time.Now().Format(time.RFC3339)+termColorReset, kind)
			fmt.Println(s.msg...)
		}
		fmt.Printf("  %d\tfunc: %s\n\tat %s:%d\n", i, s.fn, s.file, s.line)
	}
}
