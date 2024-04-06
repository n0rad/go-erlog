package erlog

import (
	"github.com/n0rad/go-erlog/errs"
	"github.com/n0rad/go-erlog/logs"
	"runtime"
	"strings"
	"time"
)

type LogEvent struct {
	logs.Entry
	Depth int              `json:"depth,omitempty"`
	Time  time.Time        `json:"time,omitempty"`
	File  string           `json:"file,omitempty"`
	Line  int              `json:"line,omitempty"`
	Err   *errs.EntryError `json:"err,omitempty"`
}

func NewLogEvent(entry *logs.Entry) *LogEvent {
	var file string
	var line int
	var ok bool
	for i := 2; ; i++ {
		_, file, line, ok = runtime.Caller(i)
		if !ok {
			file = "???"
			line = 0
		}
		if !strings.Contains(file, "n0rad/go-erlog") {
			break
		}
		if strings.Contains(file, "n0rad/go-erlog/examples") { // TODO what to do with that ?
			break
		}
	}

	return &LogEvent{
		Entry: *entry,
		File:  file,
		Line:  line,
	}
}
