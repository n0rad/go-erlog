package erlog

import (
	"encoding/json"
	"github.com/n0rad/go-erlog/data"
	"github.com/n0rad/go-erlog/errs"
	"github.com/n0rad/go-erlog/logs"
	"io"
	"sync"
	"time"
)

type ErlogJsonWriterAppender struct {
	Out   io.Writer
	Level logs.Level
	mu    sync.Mutex
}

func NewJsonErlogWriterAppender(writer io.Writer) (f *ErlogJsonWriterAppender) {
	return &ErlogJsonWriterAppender{
		Out: writer,
	}
}

func (f *ErlogJsonWriterAppender) GetLevel() logs.Level {
	return f.Level
}

func (f *ErlogJsonWriterAppender) SetLevel(level logs.Level) {
	f.Level = level
}

func (f *ErlogJsonWriterAppender) Fire(event *LogEvent) {
	event.Time = time.Now()
	switch t := event.Entry.Err.(type) {
	case nil:
	case *errs.EntryError:
		event.Err = t
	default:
		event.Err = &errs.EntryError{Message: event.Err.Error()}
	}

	jsonEvent, err := json.Marshal(event)
	if err != nil {
		jsonEvent, _ = json.Marshal(LogEvent{
			Entry: logs.Entry{
				Message: "Failed to marshal log to json",
				Fields:  data.WithField("message", event.Message),
			}})
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	_, _ = f.Out.Write(jsonEvent)
	_, _ = f.Out.Write([]byte{'\n'})
}
