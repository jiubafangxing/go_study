package main
import (

"fmt"
)

func test()(z int){
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

func main(){
	result := test()
	fmt.Println("result is ", result)
}

