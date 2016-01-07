package log
import "github.com/n0rad/go-erlog/with"

type Entry struct {
	Level   Level
	Fields  with.Data
	Message string
	Error   error
}

func WithFields(fields with.Data) *Entry {
	return &Entry{
		Fields: fields,
	}
}

func (e *Entry) Info(msg string) {
	e.Level = INFO
	e.Message = msg
	GetDefaultLog().LogEntry(e)
}

//func NewEntryLog(logger *Log) *Entry {
//	return &Entry{
//		Logger: logger,
//		Depth: 2,
//	}
//}
//
