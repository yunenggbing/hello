package main

import "fmt"

// Speaker 说话的接口
type Speaker interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}
type Pig struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + "speak!"
}
func (d Dog) Move() string {
	return d.Name + "move!"
}

func (c Cat) Speak() string {
	return c.Name + "cat speak!"
}
func (c Cat) Move() string {
	return c.Name + "move!"
}
func main() {
	d := Dog{Name: "旺财"}
	c := Cat{Name: "喵喵"}
	fmt.Println(d.Speak())
	fmt.Println(c.Speak())
	testAnimal(d)
	testAnimal(c)
}
func testAnimal(a Speaker) {
	if dog, ok := a.(Dog); ok {
		fmt.Println("this is a dog,dog name is :", dog.Name)
	} else if cat, ok := a.(Cat); ok {
		fmt.Println("this is a cat,cat name is :", cat.Name)
	} else {
		panic("unsupport type")
	}
}
