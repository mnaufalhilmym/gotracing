package gotracing

import (
	"fmt"
	"time"
)

func (st Stacktraces) String() string {
	var str string
	for i, s := range st {
		if s.msg != nil {
			str += fmt.Sprintf("%s  %s: ", termColorComment+time.Now().Format(time.RFC3339)+termColorReset, s.kind)
			for i, msg := range s.msg {
				if i > 0 {
					str += " "
				}
				str += fmt.Sprint(msg)
			}
		}
		if s.fn != "" {
			str += fmt.Sprintf("\n  %d\tfunc: %s\n\tat %s:%d", i, s.fn, s.file, s.line)
		}
	}
	return str
}
