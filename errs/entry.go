package errs

import (
	"bytes"
	"fmt"
	"github.com/n0rad/go-erlog/log"
	"github.com/n0rad/go-erlog/with"
	"runtime"
)

var MaxStackDepth = 50

type EntryError struct {
	Fields with.Data
	Msg    string
	Source error
	stack  []uintptr
	frames []StackFrame
}

func From(err error) *EntryError {
	return &EntryError{
		Source: err,
	}
}

func WithField(name string, value interface{}) *EntryError {
	return &EntryError{
		Fields: with.Field(name, value),
	}
}

func Fill(entry *EntryError) *EntryError {
	stack := make([]uintptr, MaxStackDepth)
	length := runtime.Callers(2, stack[:])
	entry.stack = stack[:length]
	return entry
}

func New(message string) *EntryError {
	entry := &EntryError{
		Msg: message,
	}
	return entry
}

func (e *EntryError) WithFields(data with.Data) *EntryError {
	e.Fields = data
	return e
}

func (e *EntryError) WithField(name string, value interface{}) *EntryError {
	if e.Fields == nil {
		e.Fields = with.Field(name, value)
	} else {
		e.Fields = e.Fields.With(name, value)
	}
	return e
}

func (e *EntryError) Message(msg string) *EntryError {
	e.Msg = msg
	return e
}

func ToLog(err error) {
	if e, ok := err.(*EntryError); ok {
		log.LogEntry(&log.Entry{
			Message: e.Msg,
			Fields:  e.Fields,
			Level:   log.WARN})
		if e.Source != nil {
			ToLog(e.Source)
		}
	} else {
		log.Warn(err.Error())
	}
}

func (e *EntryError) Error() string {
	var buffer bytes.Buffer
	buffer.WriteString(e.Msg)
	if e.Fields != nil {
		for key := range e.Fields {
			buffer.WriteString(" ")
			buffer.WriteString(key)
			buffer.WriteString("=")
			buffer.WriteString(fmt.Sprintf("%+v", e.Fields[key]))
		}
	}
	buffer.WriteString("\n")
	if e.Source != nil {
		buffer.WriteString("Caused by : ")
		buffer.WriteString(e.Source.Error())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (e *EntryError) Stack() []byte {
	buf := bytes.Buffer{}

	for _, frame := range e.StackFrames() {
		buf.WriteString(frame.String())
	}

	return buf.Bytes()
}

func (e *EntryError) StackFrames() []StackFrame {
	if e.frames == nil {
		e.frames = make([]StackFrame, len(e.stack))
		for i, pc := range e.stack {
			e.frames[i] = NewStackFrame(pc)
		}
	}
	return e.frames
}

func (e *EntryError) WithErr(err error) *EntryError {
	e.Source = err
	return e
}
