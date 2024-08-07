package chapter4
import (

"fmt"
)

func DEMO_4_3_4test()(z int){
	{
		z :=1
		// 这是一个局部变量
		//return z 
		//这种做法会报错
		return 
	}
	print("11111")
	return
}

func DEMO_4_3_4main(){
	result := DEMO_4_3_4test()
	fmt.Println("result is ", result)
}

