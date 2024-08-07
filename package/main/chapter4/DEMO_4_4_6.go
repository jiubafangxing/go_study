package chapter4

import( 
	"fmt"
)
func DEMO_4_4_6test()([]func()){
	var s []func()
	for i:=0; i<10;i++{
		s =  append(s, func (){
			fmt.Println(i)		
		})


	}
	return s
}

func DEMO_4_4_6main(){
	sl := DEMO_4_4_6test()
	for _, f := range sl{
		f()
	}
}
