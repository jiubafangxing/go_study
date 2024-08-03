package main

import( 
	"fmt"
)
func test()([]func()){
	var s []func()
	for i:=0; i<10;i++{
		s =  append(s, func (){
			fmt.Println(i)		
		})


	}
	return s
}

func main(){
	sl := test()
	for _, f := range sl{
		f()
	}
}
