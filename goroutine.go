package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producder(header string,channel chan<-string){
	for  {
		channel<-fmt.Sprintf("%s:%v",header,rand.Int31())
		time.Sleep(time.Second)
	}
}



func main() {

}
