package gotracing

import (
	"context"
	"fmt"

	"github.com/mnaufalhilmym/goasync"
)

func Trace(msg ...any) {
	printStackTrace := LevelFilterTrace.ge(conf.minConsolePrintLevel)
	storeStacktrace := LevelFilterTrace.ge(conf.minStoreLevel)

	var stacktrace Stacktraces
	if printStackTrace || storeStacktrace {
		stacktrace = traceStack(conf.maxPC, Stacktraces{{level: LevelTrace, msg: msg}})
	}

	if printStackTrace {
		fmt.Println(stacktrace)
	}
	if storeStacktrace && conf.storage != nil {
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelTrace, stacktrace)
			return nil, nil
		})
	}
}

func TraceWithMaxPC(maxPC uint, msg ...any) {
	printStackTrace := LevelFilterTrace.ge(conf.minConsolePrintLevel)
	storeStacktrace := LevelFilterTrace.ge(conf.minStoreLevel)

	var stacktrace Stacktraces
	if printStackTrace || storeStacktrace {
		stacktrace = traceStack(maxPC, Stacktraces{{level: LevelTrace, msg: msg}})
	}

	if printStackTrace {
		fmt.Println(stacktrace)
	}
	if storeStacktrace && conf.storage != nil {
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelTrace, stacktrace)
			return nil, nil
		})
	}
}

func Debug(msg ...any) {
	printStackTrace := LevelFilterDebug.ge(conf.minConsolePrintLevel)
	storeStacktrace := LevelFilterDebug.ge(conf.minStoreLevel)

	var stacktrace Stacktraces
	if printStackTrace || storeStacktrace {
		stacktrace = traceStack(conf.maxPC, Stacktraces{{level: LevelDebug, msg: msg}})
	}

	if printStackTrace {
		fmt.Println(stacktrace)
	}
	if storeStacktrace && conf.storage != nil {
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelDebug, stacktrace)
			return nil, nil
		})
	}
}

func DebugWithMaxPC(maxPC uint, msg ...any) {
	printStackTrace := LevelFilterDebug.ge(conf.minConsolePrintLevel)
	storeStacktrace := LevelFilterDebug.ge(conf.minStoreLevel)

	var stacktrace Stacktraces
	if printStackTrace || storeStacktrace {
		stacktrace = traceStack(maxPC, Stacktraces{{level: LevelDebug, msg: msg}})
	}

	if printStackTrace {
		fmt.Println(stacktrace)
	}
	if storeStacktrace && conf.storage != nil {
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelDebug, stacktrace)
			return nil, nil
		})
	}
}

func Info(msg ...any) {
	printStackTrace := LevelFilterInfo.ge(conf.minConsolePrintLevel)
	storeStacktrace := LevelFilterInfo.ge(conf.minStoreLevel)

	var stacktrace Stacktraces
	if printStackTrace || storeStacktrace {
		stacktrace = traceStack(conf.maxPC, Stacktraces{{level: LevelInfo, msg: msg}})
	}

	if printStackTrace {
		fmt.Println(stacktrace)
	}
	if storeStacktrace && conf.storage != nil {
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelInfo, stacktrace)
			return nil, nil
		})
	}
}

func InfoWithMaxPC(maxPC uint, msg ...any) {
	printStackTrace := LevelFilterInfo.ge(conf.minConsolePrintLevel)
	storeStacktrace := LevelFilterInfo.ge(conf.minStoreLevel)

	var stacktrace Stacktraces
	if printStackTrace || storeStacktrace {
		stacktrace = traceStack(maxPC, Stacktraces{{level: LevelInfo, msg: msg}})
	}

	if printStackTrace {
		fmt.Println(stacktrace)
	}
	if storeStacktrace && conf.storage != nil {
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelInfo, stacktrace)
			return nil, nil
		})
	}
}

func Warn(msg ...any) {
	printStackTrace := LevelFilterWarn.ge(conf.minConsolePrintLevel)
	storeStacktrace := LevelFilterWarn.ge(conf.minStoreLevel)

	var stacktrace Stacktraces
	if printStackTrace || storeStacktrace {
		stacktrace = traceStack(conf.maxPC, Stacktraces{{level: LevelWarn, msg: msg}})
	}

	if printStackTrace {
		fmt.Println(stacktrace)
	}
	if storeStacktrace && conf.storage != nil {
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelWarn, stacktrace)
			return nil, nil
		})
	}
}

func WarnWithMaxPC(maxPC uint, msg ...any) {
	printStackTrace := LevelFilterWarn.ge(conf.minConsolePrintLevel)
	storeStacktrace := LevelFilterWarn.ge(conf.minStoreLevel)

	var stacktrace Stacktraces
	if printStackTrace || storeStacktrace {
		stacktrace = traceStack(maxPC, Stacktraces{{level: LevelWarn, msg: msg}})
	}

	if printStackTrace {
		fmt.Println(stacktrace)
	}
	if storeStacktrace && conf.storage != nil {
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelWarn, stacktrace)
			return nil, nil
		})
	}
}

func Error(msg ...any) {
	printStackTrace := LevelFilterError.ge(conf.minConsolePrintLevel)
	storeStacktrace := LevelFilterError.ge(conf.minStoreLevel)

	var stacktrace Stacktraces
	if printStackTrace || storeStacktrace {
		stacktrace = traceStack(conf.maxPC, Stacktraces{{level: LevelError, msg: msg}})
	}

	if printStackTrace {
		fmt.Println(stacktrace)
	}
	if storeStacktrace && conf.storage != nil {
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelError, stacktrace)
			return nil, nil
		})
	}
}

func ErrorWithMaxPC(maxPC uint, msg ...any) {
	printStackTrace := LevelFilterError.ge(conf.minConsolePrintLevel)
	storeStacktrace := LevelFilterError.ge(conf.minStoreLevel)

	var stacktrace Stacktraces
	if printStackTrace || storeStacktrace {
		stacktrace = traceStack(conf.maxPC, Stacktraces{{level: LevelError, msg: msg}})
	}

	if printStackTrace {
		fmt.Println(stacktrace)
	}
	if storeStacktrace && conf.storage != nil {
		goasync.Spawn(func(ctx context.Context) (any, error) {
			conf.storage.Insert(LevelError, stacktrace)
			return nil, nil
		})
	}
}
