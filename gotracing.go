package gotracing

func Trace(msg ...any) {
	if !levelFilterTrace.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(conf.maxPC, msg...)
	printStack(levelTrace, stacktrace)
}

func TraceWithMaxPC(maxPC uint, msg ...any) {
	if !levelFilterTrace.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(maxPC, msg...)
	printStack(levelTrace, stacktrace)
}

func Debug(msg ...any) {
	if !levelFilterDebug.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(conf.maxPC, msg...)
	printStack(levelDebug, stacktrace)
}

func DebugWithMaxPC(maxPC uint, msg ...any) {
	if !levelFilterDebug.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(maxPC, msg...)
	printStack(levelDebug, stacktrace)
}

func Info(msg ...any) {
	if !levelFilterDebug.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(conf.maxPC, msg...)
	printStack(levelInfo, stacktrace)
}

func InfoWithMaxPC(maxPC uint, msg ...any) {
	if !levelFilterDebug.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(maxPC, msg...)
	printStack(levelInfo, stacktrace)
}

func Warn(msg ...any) {
	if !levelFilterDebug.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(conf.maxPC, msg...)
	printStack(levelWarn, stacktrace)
}

func WarnWithMaxPC(maxPC uint, msg ...any) {
	if !levelFilterDebug.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(maxPC, msg...)
	printStack(levelWarn, stacktrace)
}

func Error(msg ...any) {
	if !levelFilterDebug.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(conf.maxPC, msg...)
	printStack(levelError, stacktrace)
}

func ErrorWithMaxPC(maxPC uint, msg ...any) {
	if !levelFilterDebug.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(maxPC, msg...)
	printStack(levelError, stacktrace)
}
