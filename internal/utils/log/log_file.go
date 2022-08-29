package log

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type logFile struct {
	mu          sync.Mutex
	logDir      string
	logSubDir   string
	logNumber   int   // 当天的日志编号
	logSize     int64 // 单个日志大小
	projectPath string
	logFd       *os.File
}

func (f *logFile) write(prefix string, s string) {
	now := time.Now()

	// 检查子文件夹
	f.logSubDir = getSubDir(now)
	os.MkdirAll(path.Join(f.logDir, f.logSubDir), os.ModePerm)

	// 获取文件信息，没有则创建
	filePath := path.Join(f.logDir, f.logSubDir, fmt.Sprintf("%04d", f.logNumber))
	f.open(filePath)

	// 检查文件大小
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("[system] IoWrite stat err: %v\n", err)
	}
	if fileInfo.Size() > f.logSize {
		f.logNumber++
		f.open(path.Join(f.logDir, f.logSubDir, fmt.Sprintf("%04d", f.logNumber)))
	}

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	dir, _ := filepath.Split(f.projectPath)
	file = strings.TrimLeft(file, dir)
	fmt.Fprintf(f.logFd, "%s[%s] [%s:%d] %s\n", prefix, now.Format("15:04:05"), file, line, s)
}

// 子路径：202208，年+月
func getSubDir(now time.Time) string {
	return fmt.Sprintf("%04d%02d/%02d", now.Year(), now.Month(), now.Day())
}

// 按路径打开新文件
func (f *logFile) open(path string) {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		f.logFd = nil
		fmt.Printf("[system] open logfile err: %v", err)
		return
	}
	if f.logFd != nil {
		f.logFd.Sync()
		f.logFd.Close()
	}
	f.logFd = fd
}

// 设置日志输出路径
func (f *logFile) setLogDir(dir string) {
	// 检查dir是否存在
	stat, err := os.Stat(dir)
	if err != nil {
		log.Fatalf("[system] LogDir not exist err: %v\n", err)
	}
	if !stat.IsDir() {
		log.Fatalf("[system] LogDir not exist err: %v\n", err)
	}

	// 查找最新的logNumber
	now := time.Now()
	subDir := getSubDir(now)
	os.MkdirAll(path.Join(dir, subDir), os.ModePerm)
	files, err := ioutil.ReadDir(path.Join(dir, subDir))
	if err != nil {
		log.Fatalf("[system] newLogFile ReadDir err: %v\n", err)
	}

	number := 0
	if len(files) != 0 {
		number, err = strconv.Atoi(files[len(files)-1].Name())
		if err != nil {
			log.Fatalf("[system] newLogFile Atoi err: %v\n", err)
		}
	}

	f.logDir = dir
	f.logNumber = number
}

// 设置单个文件大小
func (f *logFile) setLogSize(size int64) {
	f.logSize = size
}

func (f *logFile) SetProjectPath(path string) {
	f.projectPath = path
}
