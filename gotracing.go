package gotracing

func Trace(msg string) {
	if !levelFilterTrace.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(msg)
	printStack(levelTrace, stacktrace)
}

func Debug(msg string) {
	if !levelFilterDebug.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(msg)
	printStack(levelDebug, stacktrace)
}

func Info(msg string) {
	if !levelFilterDebug.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(msg)
	printStack(levelInfo, stacktrace)
}

func Warn(msg string) {
	if !levelFilterDebug.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(msg)
	printStack(levelWarn, stacktrace)
}

func Error(msg string) {
	if !levelFilterDebug.le(conf.maxLevel) {
		return
	}
	stacktrace := traceStack(msg)
	printStack(levelError, stacktrace)
}
