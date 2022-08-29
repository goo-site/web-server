package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type logConsole struct {
	mu          sync.Mutex
	color       StringColor
	projectPath string
}

func (c *logConsole) SetProjectPath(path string) {
	c.projectPath = path
}

func (c *logConsole) write(prefix string, s string) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	// 日志信息输出到控制台
	dir, _ := filepath.Split(c.projectPath)
	file = strings.TrimLeft(file, dir)
	s1 := fmt.Sprintf("\033[%d;%dm%s[%s]\033[0m", HighLight, c.color, prefix, time.Now().Format("15:04:05"))
	s2 := fmt.Sprintf("\033[%d;%d;%dm[%s:%d]\033[0m", HighLight, UnderLine, c.color, file, line)
	s3 := fmt.Sprintf("\033[%d;%dm%s\033[0m", HighLight, c.color, s)
	fmt.Printf("%s %s %s\n", s1, s2, s3)
}

/* 输出颜色 */
type StringColor int

const (
	Red     StringColor = 31 // 红色
	Green   StringColor = 32 // 绿色
	Yellow  StringColor = 33 // 黄色
	Blue    StringColor = 34 // 蓝色
	Magenta StringColor = 35 // 洋红色
	Cyan    StringColor = 36 // 青蓝色
	White   StringColor = 37 // 白色
)

func (c *logConsole) setColor(color StringColor) {
	c.color = color
}

/* 输出样式 */
type StringAttr int

const (
	Default   StringAttr = 0 // 默认设置
	HighLight StringAttr = 1 // 高亮
	UnderLine StringAttr = 4 // 下划线
)
