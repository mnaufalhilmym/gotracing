package gotracing

import (
	"context"
	"fmt"

	"github.com/mnaufalhilmym/goasync"
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelTrace, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelTrace, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelTrace, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelTrace, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelDebug, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelDebug, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelDebug, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelDebug, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelInfo, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelInfo, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelInfo, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelInfo, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelWarn, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelWarn, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelWarn, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelWarn, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelError, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelError, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelError, traces)
			return nil, nil
		})
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
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelError, traces)
			return nil, nil
		})
	}
}
