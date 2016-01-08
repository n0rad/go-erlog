package erlog

import (
	"github.com/n0rad/go-erlog/log"
	"runtime"
	"time"
)

type LogEvent struct {
	log.Entry
	Depth int
	Time  time.Time
	File  string
	Line  int
}

func NewLogEvent(entry *log.Entry) *LogEvent {
	_, file, line, _ := runtime.Caller(4)
	return &LogEvent{
		Entry: *entry,
		File:  file,
		Line:  line,
	}
}
