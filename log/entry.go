package log

import (
	"github.com/n0rad/go-erlog/with"
)

type Entry struct {
	Logger  Log
	Level   Level
	Fields  with.Data
	Message string
	Err     error
}

func WithF(fields with.Data) *Entry {
	return &Entry{
		Logger: GetDefaultLog(),
		Fields: fields,
	}
}

func WithE(err error) *Entry {
	return &Entry{
		Logger: GetDefaultLog(),
		Err: err,
	}
}

func WithEF(err error, fields with.Data) *Entry {
	return &Entry{
		Logger: GetDefaultLog(),
		Err: err,
		Fields: fields,
	}
}

///////////////////////////////////

func (e *Entry) WithFields(data with.Data) *Entry {
	e.Fields = data
	return e
}

func (e *Entry) WithField(name string, value interface{}) *Entry {
	if e.Fields == nil {
		e.Fields = with.Field(name, value)
	} else {
		e.Fields = e.Fields.With(name, value)
	}
	return e
}

func (e *Entry) WithLog(logger Log) *Entry {
	e.Logger = logger
	return e
}

func (e *Entry) Trace(msg string) {
	e.Level = TRACE
	e.Message = msg
	e.Logger.LogEntry(e)
}

func (e *Entry) Debug(msg string) {
	e.Level = DEBUG
	e.Message = msg
	e.Logger.LogEntry(e)
}

func (e *Entry) Info(msg string) {
	e.Level = INFO
	e.Message = msg
	e.Logger.LogEntry(e)
}

func (e *Entry) Warn(msg string) {
	e.Level = WARN
	e.Message = msg
	e.Logger.LogEntry(e)
}

func (e *Entry) Error(msg string) {
	e.Level = ERROR
	e.Message = msg
	e.Logger.LogEntry(e)
}

func (e *Entry) Panic(msg string) {
	e.Level = PANIC
	e.Message = msg
	e.Logger.LogEntry(e)
}

func (e *Entry) Fatal(msg string) {
	e.Level = FATAL
	e.Message = msg
	e.Logger.LogEntry(e)
}
