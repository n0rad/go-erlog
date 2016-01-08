package erlog

import (
	"github.com/n0rad/go-erlog/log"
)

type Appender interface {
	Fire(event *LogEvent)
	GetLevel() log.Level
	SetLevel(level log.Level)
}
