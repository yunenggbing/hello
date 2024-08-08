package log

import (
	"fmt"
	"log"
	"os"
	"sync"
)

/*
@Time : 2024/8/5 17:35
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

type Logger struct {
	mu       sync.Mutex
	logLevel LogLevel
	logger   *log.Logger
}

func NewLogger(level LogLevel, output string) *Logger {
	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件：%v", err)
	}
	return &Logger{
		logLevel: level,
		logger:   log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// logMessage 根据日志级别输出日志
func (l *Logger) logMessage(level LogLevel, msg string, args ...interface{}) {
	if level < l.logLevel {
		return
	}
	message := fmt.Sprintf(msg, args...)
	switch level {
	case DEBUG:
		l.logger.Printf("[DEBUG] %s", message)
	case INFO:
		l.logger.Printf("[INFO] %s", message)
	case WARNING:
		l.logger.Printf("[WARNING] %s", message)
	case ERROR:
		l.logger.Printf("[ERROR] %s", message)
	case FATAL:
		l.logger.Fatalf("[FATAL] %s", message)
	}
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logMessage(DEBUG, msg, args...)
}
func (l *Logger) Info(msg string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logMessage(INFO, msg, args...)
}
func (l *Logger) Warning(msg string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logMessage(WARNING, msg, args...)
}
func (l *Logger) Error(msg string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logMessage(ERROR, msg, args...)
}
func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logMessage(FATAL, msg, args...)
}
