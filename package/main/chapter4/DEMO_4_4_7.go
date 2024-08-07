package chapter4

import (
	"fmt"
)
func DEMO_4_4_7test()([]func()){
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

func DEMO_4_4_7main(){
	arr := DEMO_4_4_7test()
	for _, f := range arr{
		 f()
	}

}
