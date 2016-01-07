package erlog
import (
	"os"
	"github.com/n0rad/go-erlog/log"
)


type ErrLogsFactory struct {}

func (*ErrLogsFactory) GetLog(name string) log.Log {
	return NewLog()
}

type ErrLogsLogger struct {
	appenders []Appender
	Level     log.Level
}

func NewLog() *ErrLogsLogger {
	return &ErrLogsLogger{
		appenders: []Appender{NewErlogWriterAppender(os.Stderr)},
		Level:     log.INFO,
	}
}

func (l ErrLogsLogger) log(event *LogEvent) {
	for _, appender := range l.appenders {
		appender.Fire(event)
	}
}

func (l ErrLogsLogger) Trace(message string) {
	if l.Level.IsEnableFor(log.TRACE) {
		l.log(NewLogEvent(&log.Entry{Level: log.TRACE, Message: message}))
	}
}

func (l ErrLogsLogger) Debug(message string) {
	if l.Level.IsEnableFor(log.DEBUG) {
		l.log(NewLogEvent(&log.Entry{Level: log.DEBUG, Message: message}))
	}
}
func (l ErrLogsLogger) Info(message string) {
	if l.Level.IsEnableFor(log.INFO) {
		l.log(NewLogEvent(&log.Entry{Level: log.INFO, Message: message}))
	}
}
func (l ErrLogsLogger) Warn(message string) {
	if l.Level.IsEnableFor(log.WARN) {
		l.log(NewLogEvent(&log.Entry{Level: log.WARN, Message: message}))
	}
}
func (l ErrLogsLogger) Error(message string) {
	if l.Level.IsEnableFor(log.ERROR) {
		l.log(NewLogEvent(&log.Entry{Level: log.ERROR, Message: message}))
	}
}
func (l ErrLogsLogger) Panic(message string) {
	if l.Level.IsEnableFor(log.PANIC) {
		l.log(NewLogEvent(&log.Entry{Level: log.PANIC, Message: message}))
	}
	panic(message)
}
func (l ErrLogsLogger) Fatal(message string) {
	if l.Level.IsEnableFor(log.FATAL) {
		l.log(NewLogEvent(&log.Entry{Level: log.FATAL, Message: message}))
	}
	os.Exit(1)
}
func (l ErrLogsLogger) LogEntry(entry *log.Entry) {
	if l.Level.IsEnableFor(entry.Level) {
		l.log(NewLogEvent(entry))
	}
	if entry.Level == log.PANIC {
		panic(entry.Message)
	} else if entry.Level == log.FATAL {
		os.Exit(1)
	}
}

func (l ErrLogsLogger) GetLevel() log.Level {
	return l.Level
}

func (l ErrLogsLogger) SetLevel(level log.Level) {
	l.Level = level // TODO this will not work
}

func (l ErrLogsLogger) IsTraceEnabled() bool { return log.TRACE.IsEnableFor(l.Level)}
func (l ErrLogsLogger) IsDebugEnabled() bool { return log.DEBUG.IsEnableFor(l.Level)}
func (l ErrLogsLogger) IsInfoEnabled() bool { return log.INFO.IsEnableFor(l.Level)}
func (l ErrLogsLogger) IsWarnEnabled() bool { return log.WARN.IsEnableFor(l.Level)}
func (l ErrLogsLogger) IsErrorEnabled() bool { return log.ERROR.IsEnableFor(l.Level)}
func (l ErrLogsLogger) IsPanicEnabled() bool { return log.PANIC.IsEnableFor(l.Level)}
func (l ErrLogsLogger) IsFatalEnabled() bool { return log.FATAL.IsEnableFor(l.Level)}
func (l ErrLogsLogger) IsLevelEnabled(level log.Level) bool { return level.IsEnableFor(l.Level)}


