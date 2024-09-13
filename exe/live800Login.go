package main

import (
	"bufio"
	"fmt"
	"gopkg.in/ini.v1"
	"hello/log"
	"os"
	"os/exec"
	"strings"
)

/*
@Time : 2024/9/12 17:21
@Author : echo
@File : live800
@Software: GoLand
@Description:  动态选择live800 C客户端的登录地址，并进行登录
*/
func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败:", err)
		return
	}
	filePath := dir + "\\Config\\Config.ini"
	fileContent, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer fileContent.Close()

	// 通过ini库只能拿到未被注释的
	//cfg, err := ini.Load(filePath)
	//if err != nil {
	//	fmt.Println("加载配置文件失败:", err)
	//	return
	//}
	//section := cfg.Section("Server")
	//key := section.Key("Live800Urls.0").String()
	//fmt.Println("key:" + key)
	//fmt.Println(string(file))
	logger := log.NewLogger(log.DEBUG, "exeLog.log")
	logger.Info("创建 Logger 成功")
	var values []string
	// 读取Config.ini配置文件内容
	scanner := bufio.NewScanner(fileContent)
	inSection := false
	for scanner.Scan() {
		// 读取文件中的每一行内容
		line := scanner.Text()
		//判断是否在[ ]块内部
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			inSection = true
			continue
		}
		//在[ ] 块内部，再获取所有的Live800Urls.0的配置，包含#注释的
		if inSection && (strings.HasPrefix(line, "Live800Urls.0") || strings.HasPrefix(line, "#Live800Urls.0")) {
			//strings.SplitN 函数将字符串按照指定的分隔符进行分割，返回一个字符串切片
			values = append(values, strings.SplitN(line, "=", 2)[1])
		}
	}
	// 检查扫描过程中是否有错误发生
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}
	fmt.Println("------可选项------")
	//展示配置的所有Live800Urls.0选项
	for i, value := range values {
		fmt.Printf("%d: %s\n", i+1, value)
	}
	fmt.Println("------------")

	var choice int
	fmt.Print("请选择一个选项（输入序号）: ")
	//等待用户进行输入
	_, err = fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}
	selectOption := values[choice-1]
	logger.Info("你选择了：%s", selectOption)
	fmt.Println("你选择了:", selectOption)
	// 读取 INI 配置文件
	cfg, err := ini.Load(filePath)
	if err != nil {
		logger.Fatal("Fail to read file: %v", err)
	}
	//使用用户输入的选项进行设置ini文件
	cfg.Section("Server").Key("Live800Urls.0").SetValue(selectOption)
	exePath := dir + "\\live800.exe"
	logger.Info("exe地址：%s", exePath)
	//启动exe
	cmd := exec.Command(exePath)
	cmdErr := cmd.Start()
	if cmdErr != nil {
		logger.Info("exe启动失败：%s", cmdErr)
		return
	}
	logger.Info("exe启动成功")
}
