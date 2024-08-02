package main

import "fmt"

/**
不带缓存的channel
*/

func counter1(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {

	for val := range in {
		fmt.Println(val)
	}
}

func main() {
	naturals := make(chan int)
	squarers := make(chan int)
	go counter1(naturals)
	go squarer(squarers, naturals)
	printer(squarers)

}
