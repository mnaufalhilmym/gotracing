package gotracing_test

import (
	"testing"

	"github.com/mnaufalhilmym/gotracing"
)

func TestNewLevelFilter(t *testing.T) {
	levelFilterRes := gotracing.NewLevelFilter("error")
	if levelFilterRes.IsErr() {
		t.Error("levelFilter must be Ok<T>")
		return
	}

	levelOpt := gotracing.NewLevelFilter("error")
	if levelOpt.IsOk() {
		gotracing.SetMaxLevel(levelOpt.Unwrap())
	}
	gotracing.SetMaxProgramCounters(50)

	sampleFunc(0, "TEST")
}

func sampleFunc(x int, b string) {
	x++
	if x < 10 {
		sampleFunc(x, b)
	} else {
		gotracing.Warn(b)
	}
}
