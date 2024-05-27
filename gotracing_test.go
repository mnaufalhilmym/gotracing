package gotracing_test

import (
	"testing"

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

func TestError(t *testing.T) {
	levelFilterRes := gotracing.NewLevelFilter("error")
	if levelFilterRes.IsErr() {
		t.Error("levelFilter must be Ok<T>")
		return
	}

	gotracing.SetMaxLevel(levelFilterRes.Unwrap())
	gotracing.SetMaxProgramCounters(50)

	recursionFunc(10, gotracing.Error, "Error", "ERROR")
}

func TestWarn(t *testing.T) {
	levelFilterOpt := gotracing.NewLevelFilter("warn")
	if levelFilterOpt.IsOk() {
		gotracing.SetMaxLevel(levelFilterOpt.Unwrap())
	} else {
		t.Error("levelFilter must be Ok<T>")
	}

	gotracing.SetMaxProgramCounters(10)

	recursionFunc(10, gotracing.Warn, "Warn", "WARN")
}

func TestInfoWithPC(t *testing.T) {
	levelFilterOpt := gotracing.NewLevelFilter("info")
	if levelFilterOpt.IsOk() {
		gotracing.SetMaxLevel(levelFilterOpt.Unwrap())
	} else {
		t.Error("levelFilter must be Ok<T>")
	}

	recursionFuncMaxPC(10, gotracing.InfoWithMaxPC, 1, "Info", "INFO")
}

func TestDebugWithPC(t *testing.T) {
	levelFilterOpt := gotracing.NewLevelFilter("debug")
	if levelFilterOpt.IsOk() {
		gotracing.SetMaxLevel(levelFilterOpt.Unwrap())
	} else {
		t.Error("levelFilter must be Ok<T>")
	}

	recursionFuncMaxPC(10, gotracing.DebugWithMaxPC, 0, "Info", "INFO")
}
