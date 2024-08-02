package main

import "fmt"

/*
@Time : 2024/7/24 22:11
@Author : echo
@File : memorySync
@Software: GoLand
@Description:
*/
func main() {
	var x, y int
	go func() {
		fmt.Println("func1----")
		x = 1
		fmt.Println("y:", y)
	}()

	go func() {
		fmt.Println("func2----")
		y = 1
		fmt.Println("x:", x)
	}()

}
