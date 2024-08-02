package main

import (
	"fmt"
	"time"
)

var week time.Duration
func main() {

	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n",t.Day(),t.Month(),t.Year())
	t = time.Now().UTC();
	fmt.Println(t)

	week = 60*60*24*7*1e9
	week_from_now :=t.Add(time.Duration(week))
	fmt.Println(week_from_now)
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	fmt.Println(time.Now().UnixMicro())
}
