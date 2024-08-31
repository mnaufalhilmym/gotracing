package gotracing

import (
	"fmt"
)

// The "trace" level.
//
// Designates very low priority, often extremely verbose, information.
func Trace(msg ...any) {
	printTrace := LevelFilterTrace.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterTrace.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterTrace.ge(conf.minStackTrace)

	traces := Traces{level: LevelTrace, msg: msg}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelTrace, traces)
	}
}

// The "trace" level with message formatting.
//
// Designates very low priority, often extremely verbose, information.
func Tracef(format string, a ...any) {
	printTrace := LevelFilterTrace.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterTrace.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterTrace.ge(conf.minStackTrace)

	traces := Traces{level: LevelTrace, msg: []any{fmt.Sprintf(format, a...)}}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelTrace, traces)
	}
}

// The "trace" level with forced stack trace.
//
// Designates very low priority, often extremely verbose, information.
func TraceF(msg ...any) {
	printTrace := LevelFilterTrace.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterTrace.ge(conf.minStoreLevel)

	traces := Traces{level: LevelTrace, msg: msg}
	if printTrace || storeTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelTrace, traces)
	}
}

// The "trace" level with maxPC.
//
// Designates very low priority, often extremely verbose, information.
func TracePC(maxPC uint, msg ...any) {
	printTrace := LevelFilterTrace.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterTrace.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterTrace.ge(conf.minStackTrace)

	traces := Traces{level: LevelTrace, msg: msg}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelTrace, traces)
	}
}

// The "trace" level with maxPC and forced stack traces.
//
// Designates very low priority, often extremely verbose, information.
func TracePCF(maxPC uint, msg ...any) {
	printTrace := LevelFilterTrace.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterTrace.ge(conf.minStoreLevel)

	traces := Traces{level: LevelTrace, msg: msg}
	if printTrace || storeTrace {
		traces = traceStack(maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelTrace, traces)
	}
}

// The "debug" level.
//
// Designates lower priority information.
func Debug(msg ...any) {
	printTrace := LevelFilterDebug.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterDebug.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterDebug.ge(conf.minStackTrace)

	traces := Traces{level: LevelDebug, msg: msg}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelDebug, traces)
	}
}

// The "debug" level with message formatting.
//
// Designates lower priority information.
func Debugf(format string, a ...any) {
	printTrace := LevelFilterDebug.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterDebug.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterDebug.ge(conf.minStackTrace)

	traces := Traces{level: LevelDebug, msg: []any{fmt.Sprintf(format, a...)}}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelDebug, traces)
	}
}

// The "debug" level with forced stack trace.
//
// Designates lower priority information.
func DebugF(msg ...any) {
	printTrace := LevelFilterDebug.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterDebug.ge(conf.minStoreLevel)

	traces := Traces{level: LevelDebug, msg: msg}
	if printTrace || storeTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelDebug, traces)
	}
}

// The "debug" level with maxPC.
//
// Designates lower priority information.
func DebugPC(maxPC uint, msg ...any) {
	printTrace := LevelFilterDebug.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterDebug.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterDebug.ge(conf.minStackTrace)

	traces := Traces{level: LevelDebug, msg: msg}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelDebug, traces)
	}
}

// The "debug" level with maxPC and forced stack trace.
//
// Designates lower priority information.
func DebugPCF(maxPC uint, msg ...any) {
	printTrace := LevelFilterDebug.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterDebug.ge(conf.minStoreLevel)

	traces := Traces{level: LevelDebug, msg: msg}
	if printTrace || storeTrace {
		traces = traceStack(maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelDebug, traces)
	}
}

// The "info" level.
//
// Designates useful information.
func Info(msg ...any) {
	printTrace := LevelFilterInfo.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterInfo.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterInfo.ge(conf.minStackTrace)

	traces := Traces{level: LevelInfo, msg: msg}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelInfo, traces)
	}
}

// The "info" level with message formatting.
//
// Designates useful information.
func Infof(format string, a ...any) {
	printTrace := LevelFilterInfo.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterInfo.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterInfo.ge(conf.minStackTrace)

	traces := Traces{level: LevelInfo, msg: []any{fmt.Sprintf(format, a...)}}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelInfo, traces)
	}
}

// The "info" level with forced stack trace.
//
// Designates useful information.
func InfoF(msg ...any) {
	printTrace := LevelFilterInfo.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterInfo.ge(conf.minStoreLevel)

	traces := Traces{level: LevelInfo, msg: msg}
	if printTrace || storeTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelInfo, traces)
	}
}

// The "info" level with maxPC.
//
// Designates useful information.
func InfoPC(maxPC uint, msg ...any) {
	printTrace := LevelFilterInfo.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterInfo.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterInfo.ge(conf.minStackTrace)

	traces := Traces{level: LevelInfo, msg: msg}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelInfo, traces)
	}
}

// The "info" level with maxPC and forced stack trace.
//
// Designates useful information.
func InfoPCF(maxPC uint, msg ...any) {
	printTrace := LevelFilterInfo.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterInfo.ge(conf.minStoreLevel)

	traces := Traces{level: LevelInfo, msg: msg}
	if printTrace || storeTrace {
		traces = traceStack(maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelInfo, traces)
	}
}

// The "warn" level.
//
// Designates hazardous situations.
func Warn(msg ...any) {
	printTrace := LevelFilterWarn.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterWarn.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterWarn.ge(conf.minStackTrace)

	traces := Traces{level: LevelWarn, msg: msg}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelWarn, traces)
	}
}

// The "warn" level with message formatting.
//
// Designates hazardous situations.
func Warnf(format string, a ...any) {
	printTrace := LevelFilterWarn.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterWarn.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterWarn.ge(conf.minStackTrace)

	traces := Traces{level: LevelWarn, msg: []any{fmt.Sprintf(format, a...)}}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelWarn, traces)
	}
}

// The "warn" level with forced stack trace.
//
// Designates hazardous situations.
func WarnF(msg ...any) {
	printTrace := LevelFilterWarn.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterWarn.ge(conf.minStoreLevel)

	traces := Traces{level: LevelWarn, msg: msg}
	if printTrace || storeTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelWarn, traces)
	}
}

// The "warn" level with maxPC.
//
// Designates hazardous situations.
func WarnPC(maxPC uint, msg ...any) {
	printTrace := LevelFilterWarn.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterWarn.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterWarn.ge(conf.minStackTrace)

	traces := Traces{level: LevelWarn, msg: msg}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelWarn, traces)
	}
}

// The "warn" level with maxPC and forced stack trace.
//
// Designates hazardous situations.
func WarnPCF(maxPC uint, msg ...any) {
	printTrace := LevelFilterWarn.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterWarn.ge(conf.minStoreLevel)

	traces := Traces{level: LevelWarn, msg: msg}
	if printTrace || storeTrace {
		traces = traceStack(maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelWarn, traces)
	}
}

// The "error" level.
//
// Designates very serious errors.
func Error(msg ...any) {
	printTrace := LevelFilterError.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterError.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterError.ge(conf.minStackTrace)

	traces := Traces{level: LevelError, msg: msg}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelError, traces)
	}
}

// The "error" level with message formatting.
//
// Designates very serious errors.
func Errorf(format string, a ...any) {
	printTrace := LevelFilterError.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterError.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterError.ge(conf.minStackTrace)

	traces := Traces{level: LevelError, msg: []any{fmt.Sprintf(format, a...)}}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelError, traces)
	}
}

// The "error" level with forced stack trace.
//
// Designates very serious errors.
func ErrorF(msg ...any) {
	printTrace := LevelFilterError.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterError.ge(conf.minStoreLevel)

	traces := Traces{level: LevelError, msg: msg}
	if printTrace || storeTrace {
		traces = traceStack(conf.maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelError, traces)
	}
}

// The "error" level with maxPC.
//
// Designates very serious errors.
func ErrorPC(maxPC uint, msg ...any) {
	printTrace := LevelFilterError.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterError.ge(conf.minStoreLevel)
	generateStackTrace := LevelFilterError.ge(conf.minStackTrace)

	traces := Traces{level: LevelError, msg: msg}
	if (printTrace || storeTrace) && generateStackTrace {
		traces = traceStack(maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelError, traces)
	}
}

// The "error" level with maxPC and forced stack trace.
//
// Designates very serious errors.
func ErrorPCF(maxPC uint, msg ...any) {
	printTrace := LevelFilterError.ge(conf.minConsolePrintLevel)
	storeTrace := LevelFilterError.ge(conf.minStoreLevel)

	traces := Traces{level: LevelError, msg: msg}
	if printTrace || storeTrace {
		traces = traceStack(maxPC, traces)
	}

	if printTrace {
		fmt.Println(traces)
	}
	if storeTrace && conf.storage != nil {
		go conf.storage.Insert(LevelError, traces)
	}
}
