package gotracing

import (
	"errors"
	"strings"

	"github.com/mnaufalhilmym/gooption"
	"github.com/mnaufalhilmym/goresult"
)

type Level uint8

// The "trace" level.
//
// Designates very low priority, often extremely verbose, information.
const LevelTrace Level = 0

// The "debug" level.
//
// Designates lower priority information.
const LevelDebug Level = 1

// The "info" level.
//
// Designates useful information.
const LevelInfo Level = 2

// The "warn" level.
//
// Designates hazardous situations.
const LevelWarn Level = 3

// The "error" level.
//
// Designates very serious errors.
const LevelError Level = 4

type LevelFilter struct{ gooption.Option[Level] }

func (f LevelFilter) ge(filter LevelFilter) bool {
	if f.IsSome() && filter.IsSome() {
		return f.Unwrap() >= filter.Unwrap()
	}
	return false
}

// The "off" level.
//
// Designates that trace instrumentation should be completely disabled.
var LevelFilterOff = LevelFilter{gooption.None[Level]()}

// The "error" level.
//
// Designates very serious errors.
var LevelFilterError = LevelFilter{gooption.Some(LevelError)}

// The "warn" level.
//
// Designates hazardous situations.
var LevelFilterWarn = LevelFilter{gooption.Some(LevelWarn)}

// The "info" level.
//
// Designates useful information.
var LevelFilterInfo = LevelFilter{gooption.Some(LevelInfo)}

// The "debug" level.
//
//	Designates lower priority information.
var LevelFilterDebug = LevelFilter{gooption.Some(LevelDebug)}

// The "trace" level.
//
// Designates very low priority, often extremely verbose, information.
var LevelFilterTrace = LevelFilter{gooption.Some(LevelTrace)}

// Parses a string filter to return a value of levelFilter.
//
// String filter is expected one of "off", "error", "warn", "info", "debug", "trace", or a number 0-5
func NewLevelFilter(filter string) goresult.Result[LevelFilter] {
	switch strings.ToLower(filter) {
	case "0", "off":
		return goresult.Ok(LevelFilterOff)
	case "", "1", "error":
		return goresult.Ok(LevelFilterError)
	case "2", "warn":
		return goresult.Ok(LevelFilterWarn)
	case "3", "info":
		return goresult.Ok(LevelFilterInfo)
	case "4", "debug":
		return goresult.Ok(LevelFilterDebug)
	case "5", "trace":
		return goresult.Ok(LevelFilterTrace)
	default:
		return goresult.Err[LevelFilter](errors.New("error parsing level filter: expected one of \"off\", \"error\", \"warn\", \"info\", \"debug\", \"trace\", or a number 0-5"))
	}
}
