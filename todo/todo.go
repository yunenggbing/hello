package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
@Time : 2024/8/12 15:05
@Author : echo
@File : todoList
@Software: GoLand
@Description:
*/
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TodoList struct {
	Todos []Todo `json:"todos"`
}

// 加载待办事项
func Load(fileName string) (TodoList, error) {
	var t TodoList
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			// 文件不存在，返回nil
			return t, nil
		}
		return t, err
	}
	err = json.Unmarshal(file, &t)
	return t, err
}

// 保存待办事项
func SaveTodo(fileName string, t TodoList) error {
	data, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, 0644)
}

// 添加待办事项
func (t *TodoList) Add(title string) {
	newTodo := Todo{ID: len(t.Todos) + 1, Title: title, Done: false}

	t.Todos = append(t.Todos, newTodo)
}

// 删除待办事项
func (t *TodoList) Delete(index int) {
	t.Todos = append(t.Todos[:index], t.Todos[index+1:]...)
}

// 列出待办事项
func (t *TodoList) List() {
	for i, todo := range t.Todos {
		status := " "
		if todo.Done {
			status = "✓"
			fmt.Printf("%d. %s %s\n", i+1, todo.Title, status)
		}
	}
}

// 标记待办事项为完成
func (t *TodoList) MarkDone(id int) {
	for i, todo := range t.Todos {
		if todo.ID == id {
			t.Todos[i].Done = true
		}
	}
}
