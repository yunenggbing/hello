package main

import (
	"fmt"
	"hello/newlog"
	"os"
	"time"
)

/*
@Time : 2024/8/6 15:22
@Author : echo
@File : logTest
@Software: GoLand
@Description:
*/
func main() {
	//1.0版本  log.log
	//log := log.NewLogger(log.DEBUG, "app.log")
	//log.Info("应用程序启动")
	//log.Debug("调试信息: %s", "这是一个调试消息")
	//log.Warning("警告信息: %s", "这是一个警告")
	//log.Error("错误信息: %s", "这是一个错误")
	//log.Fatal("致命错误信息: %s", "严重错误，程序将退出")

	//2.0版本  log.NewLogger
	maxSize := int64(1 * 1024)
	logger, err := newlog.NewLogger(newlog.DEBUG, "newLogger.log", maxSize, 7, 24*time.Hour)
	if err != nil {
		fmt.Printf("创建 Logger 时出错: %v\n", err)
	}
	defer logger.Stop()

	logger.Info("应用程序启动")
	for i := 1; i <= 20; i++ {
		// 使用 Logger
		logger.Debug("调试信息: %s %d", "这是一个调试消息", i)
		//logger.Warning("警告信息: %s %d", "这是一个警告", i)
		//logger.Error("错误信息: %s %d", "这是一个错误", i)
		//logger.Fatal("致命错误信息: %s %d", "严重错误，程序将退出", i)
		time.Sleep(10 * time.Millisecond) // 稍微延时写入操作
	}
	// 最终的一条信息
	logger.Info("完成写日志的测试，请检查生成的日志文件。")

	// 检查生成的文件数量
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("读取当前目录出错: %v\n", err)
		return
	}

	logFileCount := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if len(file.Name()) > 9 && file.Name()[:9] == "test_log_" {
			logFileCount++
		}
	}

	fmt.Printf("生成的日志文件数量: %d\n", logFileCount)

}
