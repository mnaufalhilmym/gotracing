package gotracing

import (
	"errors"
	"strings"

	"github.com/mnaufalhilmym/gooption"
	"github.com/mnaufalhilmym/goresult"
)

type level = uint8

// The "trace" level.
//
// Designates very low priority, often extremely verbose, information.
const levelTrace level = 0

// The "debug" level.
//
// Designates lower priority information.
const levelDebug level = 1

// The "info" level.
//
// Designates useful information.
const levelInfo level = 2

// The "warn" level.
//
// Designates hazardous situations.
const levelWarn level = 3

// The "error" level.
//
// Designates very serious errors.
const levelError level = 4

type levelFilter struct{ gooption.Option[level] }

func (f levelFilter) le(filter levelFilter) bool {
	if f.IsSome() {
		fl := f.Unwrap()
		if filter.IsSome() {
			filterLevel := filter.Unwrap()
			return fl <= filterLevel
		} else {
			return false
		}
	} else {
		if filter.IsNone() {
			return true
		} else {
			return false
		}
	}
}

// The "off" level.
//
// Designates that trace instrumentation should be completely disabled.
var levelFilterOff = levelFilter{gooption.None[level]()}

// The "error" level.
//
// Designates very serious errors.
var levelFilterError = levelFilter{gooption.Some(levelError)}

// The "warn" level.
//
// Designates hazardous situations.
var levelFilterWarn = levelFilter{gooption.Some(levelWarn)}

// The "info" level.
//
// Designates useful information.
var levelFilterInfo = levelFilter{gooption.Some(levelInfo)}

// The "debug" level.
//
//	Designates lower priority information.
var levelFilterDebug = levelFilter{gooption.Some(levelDebug)}

// The "trace" level.
//
// Designates very low priority, often extremely verbose, information.
var levelFilterTrace = levelFilter{gooption.Some(levelTrace)}

// Parses a string filter to return a value of levelFilter.
//
// String filter is expected one of "off", "error", "warn", "info", "debug", "trace", or a number 0-5
func NewLevelFilter(filter string) goresult.Result[levelFilter] {
	switch strings.ToLower(filter) {
	case "0", "off":
		return goresult.Ok(levelFilterOff)
	case "", "1", "error":
		return goresult.Ok(levelFilterError)
	case "2", "warn":
		return goresult.Ok(levelFilterWarn)
	case "3", "info":
		return goresult.Ok(levelFilterInfo)
	case "4", "debug":
		return goresult.Ok(levelFilterDebug)
	case "5", "trace":
		return goresult.Ok(levelFilterTrace)
	default:
		return goresult.Err[levelFilter](errors.New("error parsing level filter: expected one of \"off\", \"error\", \"warn\", \"info\", \"debug\", \"trace\", or a number 0-5"))
	}
}
