package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) updateName(s string) {
	p.Name = s
}

func (p *Person) updateAge(age int) {
	p.Age = age
}

func main() {
	userMap := map[string]Person{
		"a1": {Name: "张三", Age: 12},
		"a2": {Name: "李四", Age: 15},
	}
	for _, v := range userMap {
		v.updateName("1232")
	}
	fmt.Println(userMap)
	fmt.Println("*********************")
	newPerson := map[string]*Person{
		"a1": &Person{"张三", 12},
		"a2": &Person{"李思", 18},
	}
	for _, v := range newPerson {
		v.updateAge(122)
	}
	fmt.Println(*newPerson["a1"], newPerson["a2"])
}
