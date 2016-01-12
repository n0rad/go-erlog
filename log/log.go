package log

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"sync"
	"strings"
)

type LogFactory interface {
	GetLog(name string) Log
}

type Log interface {
	Trace(msg ...string)
	Debug(msg ...string)
	Info(msg ...string)
	Warn(msg ...string)
	Error(msg ...string)
	Panic(msg ...string)
	Fatal(msg ...string)

	Tracef(format string, msg ...interface{})
	Debugf(format string, msg ...interface{})
	Infof(format string, msg ...interface{})
	Warnf(format string, msg ...interface{})
	Errorf(format string, msg ...interface{})
	Panicf(format string, msg ...interface{})
	Fatalf(format string, msg ...interface{})

	LogEntry(entry *Entry)

	GetLevel() Level
	SetLevel(lvl Level)

	IsLevelEnabled(lvl Level) bool
	IsTraceEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
	IsPanicEnabled() bool
	IsFatalEnabled() bool
}

var factory LogFactory = &DummyLog{}
var mu sync.Mutex

func RegisterLoggerFactory(f LogFactory) {
	mu.Lock()
	if f == factory {
		return
	}

	if _, ok := factory.(*DummyLog); !ok {
		_, file, line, _ := runtime.Caller(1)
		fmt.Fprintf(os.Stderr, "Re-Registering the logger factory : %s:%d. There is already one registered : %s\n", file, line, reflect.TypeOf(factory))
	}
	factory = f
	mu.Unlock()
}

type DummyLog struct{}

func (d *DummyLog) GetLog(name string) Log        { return d }
func (d *DummyLog) Tracef(format string, msg ...interface{}) { d.log(TRACE, fmt.Sprintf(format, msg...))}
func (d *DummyLog) Debugf(format string, msg ...interface{}) { d.log(DEBUG, fmt.Sprintf(format, msg...))}
func (d *DummyLog) Infof(format string, msg ...interface{}) { d.log(INFO, fmt.Sprintf(format, msg...))}
func (d *DummyLog) Warnf(format string, msg ...interface{}) { d.log(WARN, fmt.Sprintf(format, msg...))}
func (d *DummyLog) Errorf(format string, msg ...interface{}) { d.log(ERROR, fmt.Sprintf(format, msg...))}
func (d *DummyLog) Panicf(format string, msg ...interface{}) { d.log(PANIC, fmt.Sprintf(format, msg...))}
func (d *DummyLog) Fatalf(format string, msg ...interface{}) { d.log(FATAL, fmt.Sprintf(format, msg...))}

func (d *DummyLog) Trace(msg ...string)              { d.log(TRACE, msg...) }
func (d *DummyLog) Debug(msg ...string)              { d.log(DEBUG, msg...) }
func (d *DummyLog) Info(msg ...string)               { d.log(INFO, msg...) }
func (d *DummyLog) Warn(msg ...string)               { d.log(WARN, msg...) }
func (d *DummyLog) Error(msg ...string)              { d.log(ERROR, msg...) }
func (d *DummyLog) Panic(msg ...string)              { d.log(PANIC, msg...); panic(msg) }
func (d *DummyLog) Fatal(msg ...string)              { d.log(FATAL, msg...); os.Exit(1) }
func (d *DummyLog) log(level Level, msg ...string) { d.LogEntry(&Entry{Level: level, Message: strings.Join(msg, " ")}) }
func (d *DummyLog) LogEntry(entry *Entry)         { fmt.Printf("%s: %s\n", entry.Level, entry.Message) }
func (d *DummyLog) GetLevel() Level               { return TRACE }
func (d *DummyLog) SetLevel(lvl Level)            { d.Error("Dummy log cannot set level") }
func (d *DummyLog) IsLevelEnabled(lvl Level) bool { return true }
func (d *DummyLog) IsTraceEnabled() bool          { return true }
func (d *DummyLog) IsDebugEnabled() bool          { return true }
func (d *DummyLog) IsInfoEnabled() bool           { return true }
func (d *DummyLog) IsWarnEnabled() bool           { return true }
func (d *DummyLog) IsErrorEnabled() bool          { return true }
func (d *DummyLog) IsPanicEnabled() bool          { return true }
func (d *DummyLog) IsFatalEnabled() bool          { return true }
