package main

import (
	"fmt"
)

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {

	//a := [2]int{}
	//fmt.Printf("a: %p\n", &a)
	//
	//test(a)
	//fmt.Println(a)
	//a := [...]int{1, 2, 3}
	//fmt.Println(a)
	//for _, v := range a {
	//	fmt.Println(v)
	//}
	//c1 := sha256.Sum256([]byte("x"))
	//c2 := sha256.Sum256([]byte("X"))
	//fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	//for _, i := range c1 {
	//	fmt.Println(i)
	//}
	//s := [...]int{1, 2, 4, 5, 76}
	//revese(s[:])
	map2 := map[string]int{
		"2": 2,
		"3": 3,
	}
	fmt.Println(map2)
	map2["4"]++
	fmt.Println(map2)
	value, ok := map2["1"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("not exist")
	}
}

func revese(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}

/*var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

func main() {
	//ages := map[string]int{
	//	"a": 31,
	//	"b": 22,
	//}
	//fmt.Println(ages["a"])
	//delete(ages, "b")
	//fmt.Println(ages["b"])
	//const (
	//	n1 = iota
	//	n2 = iota
	//	n3 = 5
	//	n4 = iota
	//)
	//const n6 = iota
	//
	//fmt.Print(n1, n2, n3, n4, n6)
	//const (
	//	_  = iota
	//	KB = 1 << (10 * iota)
	//	MB = 1 << (10 * iota)
	//	GB = 1 << (10 * iota)
	//	TB = 1 << (10 * iota)
	//	PB = 1 << (10 * iota)
	//)
	//fmt.Println(KB, MB, GB, TB, PB)
	//fmt.Println("-----------------------")
	//const (
	//	a, b = iota + 1, iota + 2 //1,2
	//	c, d                      //2,3
	//	e, f                      //3,4
	//)
	//fmt.Println(a, b, c, d, e, f)
	//changeStr()
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)
	fmt.Println(len(d))
}*/

//	func changeStr() {
//		str := "hello"
//		bytes := []byte(str)
//		bytes[0] = 'H'
//		fmt.Println(string(bytes))
//
//		s2 := "博客"
//		rune2 := []rune(s2)
//		rune2[0] = '狗'
//		fmt.Println(string(rune2))
//
// }

func createMap() {
	map1 := make(map[string]int)
	map1["1"] = 1
	map1["2"] = 2
	fmt.Println(map1)

	map2 := map[string]int{
		"2": 2,
		"3": 3,
	}
	fmt.Println(map2)
}
