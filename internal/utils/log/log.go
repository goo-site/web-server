package log

import (
	"fmt"
	"os"
)

var lc logConfig

func Init() {
	lc.SetLevel(DebugLevel)
	lc.SetOutput(FileAndConsole)

	// 文件输出设置
	if lc.logfile != nil {
		lc.logfile.SetProjectPath("/home/william/go/src/WebServer")
		lc.logfile.setLogDir("/home/william/go/src/WebServer/log")
		lc.logfile.setLogSize(64)
	}

	// 控制台输出设置
	if lc.logconsole != nil {
		lc.logconsole.SetProjectPath("/home/william/go/src/WebServer")
	}
}

func Debug(format string, args ...interface{}) {
	if lc.level < InfoLevel {
		return
	}
	if lc.logconsole != nil {
		lc.logconsole.setColor(Blue)
	}
	lc.output("[DEBUG] ", fmt.Sprintf(format, args...))
}

func Info(format string, args ...interface{}) {
	if lc.level < InfoLevel {
		return
	}
	if lc.logconsole != nil {
		lc.logconsole.setColor(Green)
	}
	lc.output("[INFO] ", fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
	if lc.level < InfoLevel {
		return
	}
	if lc.logconsole != nil {
		lc.logconsole.setColor(Yellow)
	}
	lc.output("[WARN] ", fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	if lc.level < InfoLevel {
		return
	}
	if lc.logconsole != nil {
		lc.logconsole.setColor(Red)
	}
	lc.output("[ERROR] ", fmt.Sprintf(format, args...))
}

func Panic(format string, args ...interface{}) {
	if lc.level < InfoLevel {
		return
	}
	if lc.logconsole != nil {
		lc.logconsole.setColor(Red)
	}
	lc.output("[PANIC] ", fmt.Sprintf(format, args...))
	panic(fmt.Sprintf(format, args...))
}

func Fatal(format string, args ...interface{}) {
	if lc.level < InfoLevel {
		return
	}
	if lc.logconsole != nil {
		lc.logconsole.setColor(Red)
	}
	lc.output("[FATAL] ", fmt.Sprintf(format, args...))
	os.Exit(1)
}
