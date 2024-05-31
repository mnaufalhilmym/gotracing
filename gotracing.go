package gotracing

import (
	"context"
	"fmt"

	"github.com/mnaufalhilmym/goasync"
)

func Trace(msg ...any) {
	printStackTrace := conf.minConsolePrintLevel.ge(LevelFilterTrace)
	storeStacktrace := conf.minStoreLevel.ge(LevelFilterTrace)

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
	printStackTrace := conf.minConsolePrintLevel.ge(LevelFilterTrace)
	storeStacktrace := conf.minStoreLevel.ge(LevelFilterTrace)

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
	printStackTrace := conf.minConsolePrintLevel.ge(LevelFilterDebug)
	storeStacktrace := conf.minStoreLevel.ge(LevelFilterDebug)

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
	printStackTrace := conf.minConsolePrintLevel.ge(LevelFilterDebug)
	storeStacktrace := conf.minStoreLevel.ge(LevelFilterDebug)

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
	printStackTrace := conf.minConsolePrintLevel.ge(LevelFilterInfo)
	storeStacktrace := conf.minStoreLevel.ge(LevelFilterInfo)

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
	printStackTrace := conf.minConsolePrintLevel.ge(LevelFilterInfo)
	storeStacktrace := conf.minStoreLevel.ge(LevelFilterInfo)

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
	printStackTrace := conf.minConsolePrintLevel.ge(LevelFilterWarn)
	storeStacktrace := conf.minStoreLevel.ge(LevelFilterWarn)

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
	printStackTrace := conf.minConsolePrintLevel.ge(LevelFilterWarn)
	storeStacktrace := conf.minStoreLevel.ge(LevelFilterWarn)

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
	printStackTrace := conf.minConsolePrintLevel.ge(LevelFilterError)
	storeStacktrace := conf.minStoreLevel.ge(LevelFilterError)

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
	printStackTrace := conf.minConsolePrintLevel.ge(LevelFilterError)
	storeStacktrace := conf.minStoreLevel.ge(LevelFilterError)

	var stacktrace Stacktraces
	if printStackTrace || storeStacktrace {
		stacktrace = traceStack(maxPC, Stacktraces{{level: LevelError, msg: msg}})
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
