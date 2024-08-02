package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	//str := "asSASA ddd dsjkdsjs dk";
	//fmt.Println(len(str))
	//fmt.Println(utf8.RuneCountInString(str));
	//str1 := "asSASA ddd dsjkdsjsこん dk";
	//fmt.Println(len(str1))
	//fmt.Println(utf8.RuneCountInString(str1))
	//str2 := "232dfsdfsdfsdg";
	//s := "aaa";
	//fmt.Println(strings.Replace(str2,"232",s,1)	)
	for _,url:=range  os.Args[1:]{
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(os.Stderr,"fetch: %v\n",err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(response.Body)
		err = response.Body.Close()
		if err != nil {
			fmt.Println(os.Stderr,"ftech:reading %s: %v\n",url,err)
			os.Exit(1)
		}
		fmt.Printf("成功  %s",b)
	}
}
