package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

/*
@Time : 2024/7/22 11:44
@Author : echo
@File : clockwall
@Software: GoLand
@Description:
*/
func main() {
	// 检查参数数量是否小于2
	if len(os.Args) < 2 {
		// 打印使用说明并退出
		fmt.Println("使用方法: clockwall City=host:port ...")
		os.Exit(1)
	}
	// 从参数中获取位置信息
	locations := os.Args[1:]
	// 创建一个map来存储时间
	times := make(map[string]string)
	// 创建一个通道来信号位置完成
	done := make(chan bool)
	// 遍历位置信息
	for _, location := range locations {
		// 将位置信息分割为城市和地址
		parts := strings.Split(location, "=")
		// 检查位置信息是否有效
		if len(parts) != 2 {
			// 打印错误消息并退出
			fmt.Printf("无效的位置: %s\n", location)
			os.Exit(1)
		}
		// 获取城市和地址
		city, address := parts[0], parts[1]
		// 启动一个goroutine来从地址获取时间
		go func(city, address string) {
			// 连接到地址
			conn, err := net.Dial("tcp", address)
			// 检查是否有错误
			if err != nil {
				// 打印错误消息
				fmt.Fprintf(os.Stderr, "连接到 %s 时出错: %s\n", city, err)
				// 发送信号表示位置完成
				done <- true
			}
			// 完成时关闭连接
			defer conn.Close()
			// 无限循环
			for {
				// 创建一个缓冲区来存储数据
				buf := make([]byte, 1024)
				// 从连接中读取数据
				n, err := conn.Read(buf)
				// 检查是否有错误
				if err != nil {
					// 检查错误是否不是EOF
					if err != io.EOF {
						// 打印错误消息
						fmt.Fprintf(os.Stderr, "从 %s 读取时出错: %s\n", city, err)
					}
					// 发送信号表示位置完成
					done <- true
					// 从goroutine返回
					return
				}
				// 将时间存储在映射中
				times[city] = string(buf[:n-1])
			}
		}(city, address)
	}
	// 启动一个goroutine来打印时间
	go func() {
		// 无限循环
		for {
			// 检查位置是否完成
			select {
			case <-done:
				// 从goroutine返回
				return
			default:
				// 打印时间
				printTimes(times)
				// 每秒休眠一次
				time.Sleep(1 * time.Second)
			}
		}
	}()
	// 无限循环
	for {
		// 检查位置是否完成
		select {
		case <-done:
			// 从主函数返回
			return
		}
	}
}

// 打印时间的函数
func printTimes(times map[string]string) {
	// 清空屏幕
	fmt.Print("\033[H\033[2J]")
	// 遍历时间
	for city, t := range times {
		// 打印城市和时间
		fmt.Printf("%s: %s\n", city, t)
	}
}
