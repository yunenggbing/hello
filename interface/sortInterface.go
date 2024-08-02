package main

import (
	"fmt"
	"sort"
)

/*
@Time : 2024/7/19 10:40
@Author : echo
@File : sortInterface
@Software: GoLand
@Description: sort interface接口--是否为回文串
*/

// 定义一个类型IntSlice，它是一个int切片
type IntSlice []int

func (p IntSlice) Len() int { return len(p) }

func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }

func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		if s.Less(i, s.Len()-1-i) || s.Less(s.Len()-1-i, i) {
			return false
		}
	}
	return true
}
func main() {
	intSlice := IntSlice{1, 2, 3, 2, 5, 1}
	sort.Sort(intSlice)
	fmt.Println(intSlice)
	fmt.Println(IsPalindrome(intSlice))
}
