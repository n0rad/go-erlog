package erlog

import (
	"github.com/n0rad/go-erlog/log"
	"runtime"
	"time"
	"strings"
)

type LogEvent struct {
	log.Entry
	Depth int
	Time  time.Time
	File  string
	Line  int
}

func NewLogEvent(entry *log.Entry) *LogEvent {

	var file string
	var line int
	for i := 3; ; i++ {
		_, file, line, _ = runtime.Caller(i)
		if !strings.Contains(file, "n0rad/go-erlog/log") {
			break
		}
	}

	return &LogEvent{
		Entry: *entry,
		File:  file,
		Line:  line,
	}
}
