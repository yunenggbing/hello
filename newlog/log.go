package newlog

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

/*
@Time : 2024/8/8 15:52
@Author : echo
@File : log
@Software: GoLand
@Description:
*/
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var levelToString = map[LogLevel]string{
	DEBUG:   "DEBUG",
	INFO:    "INFO",
	WARNING: "WARNING",
	ERROR:   "ERROR",
	FATAL:   "FATAL",
}

type Logger struct {
	mu               sync.Mutex
	logLevel         LogLevel
	logFile          *os.File
	logFilePath      string
	maxSize          int64
	maxAge           int64 //日志文件的最大保存时间
	rotationInterval time.Duration
	stopChan         chan struct{}
}

func NewLogger(level LogLevel, filePath string, maxSize int64, maxAge int64, rotationInterval time.Duration) (*Logger, error) {
	logger := &Logger{
		logLevel:         level,
		logFilePath:      filePath,
		maxSize:          maxSize,
		maxAge:           maxAge,
		rotationInterval: rotationInterval,
		stopChan:         make(chan struct{}),
	}
	if err := logger.rotateLog(); err != nil {
		return nil, err
	}

	//启动定时清理就日志的携程
	go logger.startLogCleanUp()
	return logger, nil
}

func (l *Logger) logMessage(level LogLevel, msg string, args ...interface{}) {
	if level < l.logLevel {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	currentTIme := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("%s [%s] %s\n", currentTIme, levelToString[level], fmt.Sprintf(msg, args...))
	if err := l.checkLogSize(); err != nil {
		return
	}
	l.logFile.WriteString(logEntry)
}

// 调整日志级别
func (l *Logger) SetLOgLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logLevel = level
}

// 检查文件大小并进行分割
func (l *Logger) checkLogSize() error {
	fileInfo, err := l.logFile.Stat()
	if err != nil {
		return err
	}
	if fileInfo.Size() > l.maxSize {
		if err := l.rotateLog(); err != nil {
			return err
		}
	}
	return nil
}

// 日志分割
func (l *Logger) rotateLog() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.logFile != nil {
		l.logFile.Close()
	}

	date := time.Now().Format("2006-01-02")

	logFileName := fmt.Sprintf("%s_%s.log", l.logFilePath, date)

	// 检查文件名是否存在，直到找到唯一的文件名
	i := 1
	for fileExists(logFileName) {
		logFileName = fmt.Sprintf("%s_%s_%d.log", l.logFilePath, date, i)
		i++
	}

	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("open log file failed:%s", err)
	}
	l.logFile = logFile
	return nil
}

// 清理旧的日志文件
func (l *Logger) startLogCleanUp() {
	ticker := time.NewTicker(l.rotationInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			l.cleanOldLogs()
		case <-l.stopChan:
			return
		}
	}
}

func (l *Logger) cleanOldLogs() {
	files, err := ioutil.ReadDir(filepath.Dir(l.logFilePath))
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if time.Since(file.ModTime()).Hours() > float64(l.maxAge*24) {
			os.Remove(filepath.Join(filepath.Dir(l.logFilePath), file.Name()))
		}
	}
}
func (l *Logger) Debug(msg string, args ...interface{}) {

	l.logMessage(DEBUG, msg, args...)
}
func (l *Logger) Info(msg string, args ...interface{}) {

	l.logMessage(INFO, msg, args...)
}
func (l *Logger) Warning(msg string, args ...interface{}) {

	l.logMessage(WARNING, msg, args...)
}
func (l *Logger) Error(msg string, args ...interface{}) {

	l.logMessage(ERROR, msg, args...)
}
func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.logMessage(FATAL, msg, args...)
	//os.Exit(1) // 退出程序
}

func (l *Logger) Stop() {
	close(l.stopChan)
	l.logFile.Close()
}

// fileExists 检查文件是否存在
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
