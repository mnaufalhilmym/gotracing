package gotracing

import (
	"fmt"
	"time"
)

func (st Traces) Format(f fmt.State, c rune) {
	var stacks string
	if !f.Flag('#') && supportsColor() {
		stacks = st.formatColor()
	} else {
		stacks = st.formatColorless()
	}
	fmt.Fprintf(f, "%s", stacks)
}

func (st *Traces) formatColor() string {
	var str string

	kind := ""
	switch st.level {
	case 0: // trace
		kind += termColorPurple + "TRACE" + termColorReset
	case 1: // debug
		kind += termColorBlue + "DEBUG" + termColorReset
	case 2: // info
		kind += termColorGreen + "INFO" + termColorReset
	case 3: // warn
		kind += termColorYellow + "WARN" + termColorReset
	case 4: // error
		kind += termColorRed + "ERROR" + termColorReset
	}
	str += fmt.Sprintf("%s  %s: ", termColorComment+time.Now().Format(time.RFC3339)+termColorReset, kind)
	for i, msg := range st.msg {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(msg)
	}

	for i, s := range st.stacktraces {
		str += fmt.Sprintf("\n  %d\tfunc: %s\n\tat %s:%d", i, s.fn, s.file, s.line)
	}

	return str
}

func (st *Traces) formatColorless() string {
	var str string

	kind := ""
	switch st.level {
	case 0: // trace
		kind = "TRACE"
	case 1: // debug
		kind = "DEBUG"
	case 2: // info
		kind = "INFO"
	case 3: // warn
		kind = "WARN"
	case 4: // error
		kind = "ERROR"
	}
	str += fmt.Sprintf("%s  %s: ", time.Now().Format(time.RFC3339), kind)
	for i, msg := range st.msg {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(msg)
	}

	for i, s := range st.stacktraces {
		str += fmt.Sprintf("\n  %d\tfunc: %s\n\tat %s:%d", i, s.fn, s.file, s.line)
	}

	return str
}
