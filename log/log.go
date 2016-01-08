package log

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"sync"
)

type LogFactory interface {
	GetLog(name string) Log
}

type Log interface {
	Trace(msg string)
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Panic(msg string)
	Fatal(msg string)
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
func (d *DummyLog) Trace(msg string)              { d.LogEntry(&Entry{Level: TRACE, Message: msg}) }
func (d *DummyLog) Debug(msg string)              { d.LogEntry(&Entry{Level: DEBUG, Message: msg}) }
func (d *DummyLog) Info(msg string)               { d.LogEntry(&Entry{Level: INFO, Message: msg}) }
func (d *DummyLog) Warn(msg string)               { d.LogEntry(&Entry{Level: WARN, Message: msg}) }
func (d *DummyLog) Error(msg string)              { d.LogEntry(&Entry{Level: ERROR, Message: msg}) }
func (d *DummyLog) Panic(msg string)              { d.LogEntry(&Entry{Level: PANIC, Message: msg}); panic(msg) }
func (d *DummyLog) Fatal(msg string)              { d.LogEntry(&Entry{Level: FATAL, Message: msg}); os.Exit(1) }
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
