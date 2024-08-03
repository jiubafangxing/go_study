package main

import (
	"fmt"
)
func test()([]func()){
	var arr []func()
	for i:=0;i<=3;i++ {
		x := i
		arr = append(arr, func(){
			fmt.Println("start")
			fmt.Println(x)		
			fmt.Println("end")
		})
	}
	fmt.Println(len(arr))
	return arr
}

func main(){
	arr := test()
	for _, f := range arr{
		 f()
	}

}
