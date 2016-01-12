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
	Fields  with.Data
	Message string
	Err     error
	stack   []uintptr
	frames  []StackFrame
}

func WithSource(err error) *EntryError {
	return Fill(&EntryError{
		Err: err,
	})
}

func WithField(name string, value interface{}) *EntryError {
	return Fill(&EntryError{
		Fields: with.Field(name, value),
	})
}

func WithMessage(message string) *EntryError {
	return Fill(&EntryError{
		Message: message,
	})
}

func Fill(entry *EntryError) *EntryError {
	stack := make([]uintptr, MaxStackDepth)
	length := runtime.Callers(2, stack[:])
	entry.stack = stack[:length]
	return entry
}

func FromF(fields with.Data, msg string) *EntryError {
	return Fill(&EntryError{
		Fields:  fields,
		Message: msg,
	})
}

func FromE(err error, msg string) *EntryError {
	return Fill(&EntryError{
		Err:     err,
		Message: msg,
	})
}

func FromEF(err error, fields with.Data, msg string) *EntryError {
	return Fill(&EntryError{
		Err:     err,
		Fields:  fields,
		Message: msg,
	})
}

///////////////////////////////////////////////

func Is(e1 error, e2 error) bool {
	if e1 == e2 {
		return true
	}

	ee1, ok1 := e1.(*EntryError)
	ee2, ok2 := e2.(*EntryError)
	if ok1 && ok2 && ee1.Message == ee2.Message {
		return true
	}

	if e1.Error() == e2.Error() {
		return true
	}

	return false
}

func ToFatal(err error) { toLog(err, log.FATAL) }
func ToPanic(err error) { toLog(err, log.PANIC) }
func ToError(err error) { toLog(err, log.ERROR) }
func ToWarn(err error)  { toLog(err, log.WARN) }
func ToInfo(err error)  { toLog(err, log.INFO) }
func ToDebug(err error) { toLog(err, log.DEBUG) }
func ToTrace(err error) { toLog(err, log.TRACE) }

func toLog(err error, level log.Level) {
	if e, ok := err.(*EntryError); ok {
		log.LogEntry(&log.Entry{
			Message: e.Message,
			Fields:  e.Fields,
			Level:   level})
		if e.Err != nil { // TODO this sux
			toLog(e.Err, level)
		}
	} else {
		log.LogEntry(&log.Entry{
			Message: err.Error(),
			Level:   level,
		})
	}
}

//////////////////////////////////////////////

func (e *EntryError) WithFields(data with.Data) *EntryError {
	e.Fields = data
	return e
}

func (e *EntryError) WithErr(err error) *EntryError {
	e.Err = err
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

func (e *EntryError) WithMessage(msg string) *EntryError {
	e.Message = msg
	return e
}

func (e *EntryError) Error() string {
	var buffer bytes.Buffer
	buffer.WriteString(e.Message)
	if e.Fields != nil {
		for key := range e.Fields {
			buffer.WriteString(" ")
			buffer.WriteString(key)
			buffer.WriteString("=")
			buffer.WriteString(fmt.Sprintf("%+v", e.Fields[key]))
		}
	}
	buffer.WriteString("\n")
	if e.Err != nil {
		buffer.WriteString("Caused by : ")
		buffer.WriteString(e.Err.Error())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

//
//func (e *EntryError) Stack() []byte {
//	buf := bytes.Buffer{}
//
//	for _, frame := range e.StackFrames() {
//		buf.WriteString(frame.String())
//	}
//
//	return buf.Bytes()
//}
//
//func (e *EntryError) StackFrames() []StackFrame {
//	if e.frames == nil {
//		e.frames = make([]StackFrame, len(e.stack))
//		for i, pc := range e.stack {
//			e.frames[i] = NewStackFrame(pc)
//		}
//	}
//	return e.frames
//}
