package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"hello/log"
	"io/ioutil"
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
const fileName = "optionUrl.json"

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败:", err)
		return
	}
	//json文件路径
	jsonFileName := dir + "\\" + fileName
	fmt.Println("json文件路径:", jsonFileName)
	ops, err := Load(jsonFileName)
	if err != nil {
		fmt.Println("读取json文件失败:", err)
		return
	}
	var values []string
	logger := log.NewLogger(log.DEBUG, "exeLog.log")
	filePath := dir + "\\Config\\Config.ini"
	var scanner *bufio.Scanner
	done := false
	if len(ops.Options) == 0 {
		fmt.Println("没有可用的登录地址,重新生成")
		scanner, done, values = initOptionJSon(dir, err, logger, values, ops, jsonFileName)
		if done {
			return
		}
		// 检查扫描过程中是否有错误发生
		if err := scanner.Err(); err != nil {
			fmt.Println("读取文件失败:", err)
			return
		}
	} else {
		for _, option := range ops.Options {
			values = append(values, option.Value)
		}
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
	// 通过ini库读取 INI 配置文件
	//cfg, err := ini.Load(filePath)
	//if err != nil {
	//	logger.Fatal("Fail to read file: %v", err)
	//}
	////使用用户输入的选项进行设置ini文件
	//section := cfg.Section("Server")
	//section.Key("Live800Urls.0").SetValue(selectOption)

	//err1 := cfg.SaveTo(filePath)
	//if err1 != nil {
	//	logger.Fatal("Fail to save file: %v", err1)
	//	return
	//}
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		logger.Error("Fail to open file: %v", err)

	}
	defer file.Close()
	//使用用户输入的选项进行设置ini文件
	ResetConfigBySelectUrl(file, selectOption)

	newCfg, err2 := ini.Load(filePath)
	if err2 != nil {
		logger.Fatal("Fail to reload file:%v", err2)
	}
	logger.Info("重新获取后的地址：%s", newCfg.Section("Server").Key("Live800Urls.0"))
	exePath := dir + "\\live800.exe"
	logger.Info("exe地址：%s", exePath)
	////启动exe
	cmd := exec.Command(exePath)
	cmdErr := cmd.Start()
	if cmdErr != nil {
		logger.Info("exe启动失败：%s", cmdErr)
		return
	}
	logger.Info("exe启动成功")
}

func initOptionJSon(dir string, err error, logger *log.Logger, values []string, ops Options, jsonFileName string) (*bufio.Scanner, bool, []string) {
	filePath := dir + "\\Config\\Config.ini"
	fileContent, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return nil, true, values
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

	logger.Info("创建 Logger 成功")
	scanner, values := readIniData(fileContent, values)
	for v := range values {
		fmt.Printf("values[%d] = %s\n", v, values[v])
		//保存到option.json中
		ops.Add(values[v])
		err := SaveOption(jsonFileName, ops)
		if err != nil {
			logger.Error("保存选项失败：path:%s,value:%s ,error:%s", filePath, values[v], err)
			continue
		}
	}
	return scanner, false, values
}

func readIniData(fileContent *os.File, values []string) (*bufio.Scanner, []string) {
	// 读取Config.ini配置文件内容
	scanner := bufio.NewScanner(fileContent)
	inSection := false
	var lines []string
	for scanner.Scan() {
		// 读取文件中的每一行内容
		line := scanner.Text()
		//判断是否在[ ]块内部
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			inSection = true
			lines = append(lines, line)
			continue
		}
		newLine := strings.ReplaceAll(line, " ", "")
		//在[ ] 块内部，再获取所有的Live800Urls.0的配置，包含#注释的
		if inSection && (strings.HasPrefix(newLine, "Live800Urls.0") || strings.HasPrefix(newLine, "#Live800Urls.0")) {
			//strings.SplitN 函数将字符串按照指定的分隔符进行分割，返回一个字符串切片
			values = append(values, strings.SplitN(newLine, "=", 2)[1])
			//如果不存在#注释，则记上注释
			if !strings.HasPrefix(newLine, "#") {
				newLine = "#Live800Urls.0=" + newLine
			}
			lines = append(lines, newLine)
		} else {
			lines = append(lines, line)
		}
	}
	reSaveFile(fileContent, lines)
	return scanner, values
}

func reSaveFile(fileContent *os.File, lines []string) {
	if err := fileContent.Truncate(0); err != nil {
		fmt.Println("文件截断失败：", err)
	}
	//将文件指针移动到文件开头
	if _, err := fileContent.Seek(0, 0); err != nil {
		fmt.Println("文件指针移动失败：", err)
	}
	//将切片中的内容写入文件
	writer := bufio.NewWriter(fileContent)
	for line := range lines {
		_, err := writer.WriteString(lines[line] + "\n")
		if err != nil {
			fmt.Println(fmt.Println("写入文件失败：", err))
			continue
		}
	}
	//查看缓存中的内容是否完全写入文件中
	if err := writer.Flush(); err != nil {
		fmt.Println("写入文件失败：", err)
	}
}

type Option struct {
	Value string `json:"value"`
}

type Options struct {
	Options []Option `json:"options"`
}

func SaveOption(filePath string, ops Options) error {
	data, err := json.MarshalIndent(ops, "", "")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

// Load 从文件中加载配置
func Load(filePath string) (Options, error) {
	var op Options
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// 文件不存在，创建一个空的Options对象
			return op, nil
		}
		return op, err
	}
	err = json.Unmarshal(file, &op)
	return op, err
}

func (ops *Options) Add(url string) {
	newOp := Option{Value: url}
	ops.Options = append(ops.Options, newOp)
}

// ResetConfigBySelectUrl 根据所选择的选项打开config中对应的url配置
func ResetConfigBySelectUrl(file *os.File, selectUrl string) {
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, selectUrl) {
			line = strings.ReplaceAll(line, "#", "")
		}
		lines = append(lines, line)
	}
	reSaveFile(file, lines)

}
