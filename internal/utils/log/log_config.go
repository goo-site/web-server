package log

type logConfig struct {
	level      logLevel    // 日志等级
	logconsole *logConsole // 日志输出
	logfile    *logFile    // 文件输出
}

/* 日志级别 */
type logLevel int

const (
	PanicLevel logLevel = 0
	FatalLevel logLevel = 1
	ErrorLevel logLevel = 2
	WarnLevel  logLevel = 3
	InfoLevel  logLevel = 4
	DebugLevel logLevel = 5
)

func (l *logConfig) SetLevel(level logLevel) {
	l.level = level
}

/* 输出模式 */
type outputMode int

const (
	File           outputMode = 0
	Console        outputMode = 1
	FileAndConsole outputMode = 2
)

func (l *logConfig) SetOutput(mode outputMode) {
	if mode == File || mode == FileAndConsole {
		l.logfile = &logFile{}
	}
	if mode == Console || mode == FileAndConsole {
		l.logconsole = &logConsole{}
	}
}

func (l *logConfig) output(prefix string, s string) {
	if l.logfile != nil {
		l.logfile.write(prefix, s)
	}
	if l.logconsole != nil {
		l.logconsole.write(prefix, s)
	}
}
