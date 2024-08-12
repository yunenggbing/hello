package main

import (
	"fmt"
	"hello/todo"
	"os"
)

/*
@Time : 2024/8/12 16:28
@Author : echo
@File : todo
@Software: GoLand
@Description:
*/

const filename = "todos.json"

func main() {
	todoList, err := todo.Load(filename)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("加载待办事项失败:", err)
		return
	}
	for {
		fmt.Println("1. 添加待办事项")
		fmt.Println("2. 标记待办事项为已完成")
		fmt.Println("3. 查看待办事项")
		fmt.Println("4. 删除待办事项")
		fmt.Println("5. 退出")
		fmt.Print("选择操作: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("无效的输入，请重试。")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("输入待办事项任务: ")
			var task string
			fmt.Scan(&task)
			todoList.Add(task)
			err := todo.SaveTodo(filename, todoList)
			if err != nil {
				return
			}
			fmt.Println("已添加待办事项:", task)

		case 2:
			fmt.Print("输入待办事项 ID: ")
			var id int
			fmt.Scan(&id)
			todoList.MarkDone(id)
			todo.SaveTodo(filename, todoList)
			fmt.Println("已标记待办事项 ID 为已完成:", id)

		case 3:
			fmt.Println("待办事项列表:")
			if len(todoList.Todos) == 0 {
				fmt.Println("没有待办事项。")
			} else {
				for _, todo := range todoList.Todos {
					status := "未完成"
					if todo.Done {
						status = "已完成"
					}
					fmt.Printf("ID: %d, 任务: %s, 状态: %s\n", todo.ID, todo.Title, status)
				}
			}
		case 4:
			fmt.Print("输入待办事项 ID: ")
			if len(todoList.Todos) == 0 {
				fmt.Println("没有待办事项。")
			} else {
				var id int
				fmt.Scan(&id)
				todoList.Delete(id)
				todo.SaveTodo(filename, todoList)
				fmt.Println("已删除待办事项 ID:", id)
			}

		case 5:
			fmt.Println("再见！")
			return

		default:
			fmt.Println("无效的选择，请重试。")
		}
	}
}
