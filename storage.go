package gotracing

type storage interface {
	Insert(Level, Stacktraces)
}
