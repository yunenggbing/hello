package main

import "hello/log"

/*
@Time : 2024/8/6 15:22
@Author : echo
@File : logTest
@Software: GoLand
@Description:
*/
func main() {
	log := log.NewLogger(log.DEBUG, "app.log")
	log.Info("应用程序启动")
	log.Debug("调试信息: %s", "这是一个调试消息")
	log.Warning("警告信息: %s", "这是一个警告")
	log.Error("错误信息: %s", "这是一个错误")
	log.Fatal("致命错误信息: %s", "严重错误，程序将退出")
}
