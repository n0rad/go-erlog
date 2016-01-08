package erlog

import (
	"github.com/n0rad/go-erlog/log"
	"os"
)

type ErlogFactory struct{
	defaultLog *ErlogLogger
	logs map[string]*ErlogLogger
}

func NewErlogFactory() *ErlogFactory {
	return &ErlogFactory{
		defaultLog: newLog(),
		logs: make(map[string]*ErlogLogger, 10),
	}
}

func (l *ErlogFactory) GetLog(name string) log.Log {
	if name == "" {
		return l.defaultLog
	}
	log := l.logs[name]
	if log == nil {
		log = newLog()
		l.logs[name] = log
	}
	return log
}

type ErlogLogger struct {
	Appenders []Appender
	Level     log.Level
}

func newLog() *ErlogLogger {
	return &ErlogLogger{
		Appenders: []Appender{NewErlogWriterAppender(os.Stderr)},
		Level:     log.INFO,
	}
}

func (l *ErlogLogger) log(event *LogEvent) {
	for _, appender := range l.Appenders {
		appender.Fire(event)
	}
}

func (l *ErlogLogger) Trace(message string) {
	if log.TRACE.IsEnableFor(l.Level) {
		l.log(NewLogEvent(&log.Entry{Level: log.TRACE, Message: message}))
	}
}

func (l *ErlogLogger) Debug(message string) {
	if log.DEBUG.IsEnableFor(l.Level) {
		l.log(NewLogEvent(&log.Entry{Level: log.DEBUG, Message: message}))
	}
}
func (l *ErlogLogger) Info(message string) {
	if log.INFO.IsEnableFor(l.Level) {
		l.log(NewLogEvent(&log.Entry{Level: log.INFO, Message: message}))
	}
}
func (l *ErlogLogger) Warn(message string) {
	if log.WARN.IsEnableFor(l.Level) {
		l.log(NewLogEvent(&log.Entry{Level: log.WARN, Message: message}))
	}
}
func (l *ErlogLogger) Error(message string) {
	if log.ERROR.IsEnableFor(l.Level) {
		l.log(NewLogEvent(&log.Entry{Level: log.ERROR, Message: message}))
	}
}
func (l *ErlogLogger) Panic(message string) {
	if log.PANIC.IsEnableFor(l.Level) {
		l.log(NewLogEvent(&log.Entry{Level: log.PANIC, Message: message}))
	}
	panic(message)
}
func (l *ErlogLogger) Fatal(message string) {
	if log.FATAL.IsEnableFor(l.Level) {
		l.log(NewLogEvent(&log.Entry{Level: log.FATAL, Message: message}))
	}
	os.Exit(1)
}
func (l *ErlogLogger) LogEntry(entry *log.Entry) {
	if entry.Level.IsEnableFor(l.Level) {
		l.log(NewLogEvent(entry))
	}
	if entry.Level == log.PANIC {
		panic(entry.Message)
	} else if entry.Level == log.FATAL {
		os.Exit(1)
	}
}

func (l *ErlogLogger) GetLevel() log.Level {return l.Level}
func (l *ErlogLogger) SetLevel(level log.Level) {l.Level = level}

func (l *ErlogLogger) IsTraceEnabled() bool                { return log.TRACE.IsEnableFor(l.Level) }
func (l *ErlogLogger) IsDebugEnabled() bool                { return log.DEBUG.IsEnableFor(l.Level) }
func (l *ErlogLogger) IsInfoEnabled() bool                 { return log.INFO.IsEnableFor(l.Level) }
func (l *ErlogLogger) IsWarnEnabled() bool                 { return log.WARN.IsEnableFor(l.Level) }
func (l *ErlogLogger) IsErrorEnabled() bool                { return log.ERROR.IsEnableFor(l.Level) }
func (l *ErlogLogger) IsPanicEnabled() bool                { return log.PANIC.IsEnableFor(l.Level) }
func (l *ErlogLogger) IsFatalEnabled() bool                { return log.FATAL.IsEnableFor(l.Level) }
func (l *ErlogLogger) IsLevelEnabled(level log.Level) bool { return level.IsEnableFor(l.Level) }
