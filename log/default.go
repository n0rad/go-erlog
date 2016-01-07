package log

func GetDefaultLog() Log {return factory.GetLog("")}
func GetLog(name string) Log {return factory.GetLog(name)}

func SetLevel(lvl Level) {GetDefaultLog().SetLevel(lvl)}
func GetLevel() Level {return GetDefaultLog().GetLevel()}

func Trace(msg string) {GetDefaultLog().Trace(msg)}
func Debug(msg string) {GetDefaultLog().Debug(msg)}
func Info(msg string) {GetDefaultLog().Info(msg)}
func Warn(msg string) {GetDefaultLog().Warn(msg)}
func Error(msg string) {GetDefaultLog().Error(msg)}
func Panic(msg string) {GetDefaultLog().Panic(msg)}
func Fatal(msg string) {GetDefaultLog().Fatal(msg)}
func LogEntry(entry *Entry) {GetDefaultLog().LogEntry(entry)}

func IsLevelEnabled(lvl Level) bool {return GetDefaultLog().IsLevelEnabled(lvl)}
func IsTraceEnabled() bool {return GetDefaultLog().IsTraceEnabled()}
func IsDebugEnabled() bool {return GetDefaultLog().IsDebugEnabled()}
func IsInfoEnabled() bool {return GetDefaultLog().IsInfoEnabled()}
func IsWarnEnabled() bool {return GetDefaultLog().IsWarnEnabled()}
func IsErrorEnabled() bool {return GetDefaultLog().IsErrorEnabled()}
func IsPanicEnabled() bool {return GetDefaultLog().IsPanicEnabled()}
func IsFatalEnabled() bool {return GetDefaultLog().IsFatalEnabled()}
