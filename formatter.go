package erlog

import (
	"bytes"
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/n0rad/go-erlog/logs"
	"io"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

var pathSkip int = 0

var reset = ansi.ColorCode("reset")

var fileColorNormal = ansi.ColorCode("cyan+b")
var fileColorFail = ansi.ColorCode("cyan")

var timeColorNormal = ansi.ColorFunc("blue+b")
var timeColorFail = ansi.ColorFunc("blue")

var lvlColorError = ansi.ColorCode("red+b")
var lvlColorWarn = ansi.ColorCode("yellow+b")
var lvlColorInfo = ansi.ColorCode("green")
var lvlColorDebug = ansi.ColorCode("magenta")
var lvlColorTrace = ansi.ColorCode("blue")
var lvlColorPanic = ansi.ColorCode(":red+h")

type ErlogWriterAppender struct {
	Out   io.Writer
	Level logs.Level
	mu    sync.Mutex
}

func init() {
	_, file, _, _ := runtime.Caller(0)
	paths := strings.Split(file, "/")
	for i := 0; i < len(paths); i++ {
		if paths[i] == "github.com" {
			pathSkip = i + 2
			break
		}
	}
}

func NewErlogWriterAppender(writer io.Writer) (f *ErlogWriterAppender) {
	return &ErlogWriterAppender{
		Out: writer,
	}
}

func (f *ErlogWriterAppender) GetLevel() logs.Level {
	return f.Level
}

func (f *ErlogWriterAppender) SetLevel(level logs.Level) {
	f.Level = level
}

func (f *ErlogWriterAppender) Fire(event *LogEvent) {
	keys := f.prepareKeys(event)
	time := time.Now().Format("15:04:05")
	level := f.textLevel(event.Level)

	//	isColored := isTerminal && (runtime.GOOS != "windows")
	paths := strings.SplitN(event.File, "/", pathSkip+1)

	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%s %s%-5s%s %s%30s:%-3d%s %s%-44s%s",
		f.timeColor(event.Level)(time),
		f.levelColor(event.Level),
		level,
		reset,
		f.fileColor(event.Level),
		f.reduceFilePath(paths[pathSkip], 30),
		event.Line,
		reset,
		f.textColor(event.Level),
		event.Message,
		reset)
	for _, k := range keys {
		v := event.Entry.Fields[k]
		fmt.Fprintf(b, " %s%s%s=%+v", lvlColorInfo, k, reset, v)
	}
	b.WriteByte('\n')

	//	f.mu.Lock() //TODO
	f.Out.Write(b.Bytes())
	//	f.mu.Unlock()
}

func (f *ErlogWriterAppender) reduceFilePath(path string, max int) string {
	if len(path) <= max {
		return path
	}

	split := strings.Split(path, "/")
	splitlen := len(split)
	reducedSize := len(path)
	var buffer bytes.Buffer
	for i, e := range split {
		if reducedSize > max && i+1 < splitlen {
			buffer.WriteByte(e[0])
			reducedSize -= len(e) - 1
		} else {
			buffer.WriteString(e)
		}
		if i+1 < splitlen {
			buffer.WriteByte('/')
		}
	}
	return buffer.String()
}

func (f *ErlogWriterAppender) prepareKeys(event *LogEvent) []string {
	var keys []string = make([]string, 0, len(event.Entry.Fields))
	for k := range event.Entry.Fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (f *ErlogWriterAppender) textLevel(level logs.Level) string {
	levelText := strings.ToUpper(level.String())
	switch level {
	case logs.INFO:
	case logs.WARN:
		levelText = levelText[0:4]
	default:
		levelText = levelText[0:5]
	}
	return levelText
}

func (f *ErlogWriterAppender) fileColor(level logs.Level) string {
	switch level {
	case logs.DEBUG, logs.INFO, logs.TRACE:
		return fileColorFail
	default:
		return fileColorNormal
	}
}

func (f *ErlogWriterAppender) textColor(level logs.Level) string {
	switch level {
	case logs.WARN:
		return lvlColorWarn
	case logs.ERROR, logs.FATAL, logs.PANIC:
		return lvlColorError
	default:
		return ""
	}
}

func (f *ErlogWriterAppender) timeColor(level logs.Level) func(string) string {
	switch level {
	case logs.DEBUG, logs.INFO, logs.TRACE:
		return timeColorFail
	default:
		return timeColorNormal
	}
}

func (f *ErlogWriterAppender) levelColor(level logs.Level) string {
	switch level {
	case logs.TRACE:
		return lvlColorTrace
	case logs.DEBUG:
		return lvlColorDebug
	case logs.WARN:
		return lvlColorWarn
	case logs.ERROR:
		return lvlColorError
	case logs.FATAL, logs.PANIC:
		return lvlColorPanic
	default:
		return lvlColorInfo
	}
}

func needsQuoting(text string) bool {
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.') {
			return false
		}
	}
	return true
}

func (f *ErlogWriterAppender) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	b.WriteString(key)
	b.WriteByte('=')

	switch value := value.(type) {
	case string:
		if needsQuoting(value) {
			b.WriteString(value)
		} else {
			fmt.Fprintf(b, "%q", value)
		}
	case error:
		errmsg := value.Error()
		if needsQuoting(errmsg) {
			b.WriteString(errmsg)
		} else {
			fmt.Fprintf(b, "%q", value)
		}
	default:
		fmt.Fprint(b, value)
	}

	b.WriteByte(' ')
}
