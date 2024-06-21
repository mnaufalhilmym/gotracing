package gotracing_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mnaufalhilmym/gotracing"
)

func recursionFunc(x int, fn func(...any), msg ...any) {
	x--
	if x >= 0 {
		recursionFunc(x, fn, msg...)
	} else {
		fn(msg...)
	}
}

func recursionFuncMaxPC(x int, fn func(uint, ...any), maxPC uint, msg ...any) {
	x--
	if x >= 0 {
		recursionFuncMaxPC(x, fn, maxPC, msg...)
	} else {
		fn(maxPC, msg...)
	}
}

type storage struct{}

func (storage) Insert(l gotracing.Level, s gotracing.Traces) {
	fmt.Println("insert level", l)
	fmt.Println("insert trace", s)
}

func TestError(t *testing.T) {
	levelFilterRes := gotracing.NewLevelFilter("error")
	if levelFilterRes.IsErr() {
		t.Error("levelFilter must be Ok<T>")
		return
	}

	gotracing.SetMinConsolePrintLevel(levelFilterRes.Unwrap())
	gotracing.SetMaxProgramCounters(50)

	recursionFunc(10, gotracing.Error, "Error", "ERROR")
}

func TestWarn(t *testing.T) {
	levelFilterOpt := gotracing.NewLevelFilter("warn")
	if levelFilterOpt.IsOk() {
		gotracing.SetMinConsolePrintLevel(levelFilterOpt.Unwrap())
	} else {
		t.Error("levelFilter must be Ok<T>")
	}

	gotracing.SetMaxProgramCounters(10)

	recursionFunc(10, gotracing.Warn, "Warn", "WARN")
}

func TestInfoWithPC(t *testing.T) {
	levelFilterOpt := gotracing.NewLevelFilter("info")
	if levelFilterOpt.IsOk() {
		gotracing.SetMinConsolePrintLevel(levelFilterOpt.Unwrap())
	} else {
		t.Error("levelFilter must be Ok<T>")
	}

	recursionFuncMaxPC(10, gotracing.InfoPC, 1, "Info", "INFO")
}

func TestDebugWithPC(t *testing.T) {
	levelFilterOpt := gotracing.NewLevelFilter("debug")
	if levelFilterOpt.IsOk() {
		gotracing.SetMinConsolePrintLevel(levelFilterOpt.Unwrap())
	} else {
		t.Error("levelFilter must be Ok<T>")
	}

	recursionFuncMaxPC(10, gotracing.DebugPC, 0, "debug", "DEBUG")
}

func TestSetMinStoreLevel(t *testing.T) {
	levelFilterOpt := gotracing.NewLevelFilter("debug")
	if levelFilterOpt.IsOk() {
		gotracing.SetMinConsolePrintLevel(levelFilterOpt.Unwrap())
		gotracing.SetMinStoreLevel(levelFilterOpt.Unwrap())
	} else {
		t.Error("levelFilter must be Ok<T>")
	}

	recursionFuncMaxPC(10, gotracing.DebugPC, 10, "debug", "DEBUG")
	time.Sleep(1 * time.Second)
}

func TestStorage(t *testing.T) {
	levelFilterOpt := gotracing.NewLevelFilter("debug")
	if levelFilterOpt.IsOk() {
		gotracing.SetMinConsolePrintLevel(levelFilterOpt.Unwrap())
		gotracing.SetMinStoreLevel(levelFilterOpt.Unwrap())
	} else {
		t.Error("levelFilter must be Ok<T>")
	}
	storage := storage{}
	gotracing.SetStorage(storage)

	recursionFuncMaxPC(10, gotracing.DebugPC, 10, "debug", "DEBUG")
	time.Sleep(1 * time.Second)
}

func TestLevel(t *testing.T) {
	gotracing.SetMinConsolePrintLevel(gotracing.LevelFilterOff)
	gotracing.InfoPC(0, "test-info")
	gotracing.ErrorPC(0, "test-error")
	gotracing.SetMinConsolePrintLevel(gotracing.LevelFilterInfo)
	gotracing.TracePC(0, "test-trace")
	gotracing.InfoPC(0, "test-info2")
	gotracing.ErrorPC(0, "test-error2")
	gotracing.SetMinConsolePrintLevel(gotracing.LevelFilterError)
	gotracing.InfoPC(0, "test-info2")
	gotracing.Error("test-error3")
}
