package chapter4
import(
	"fmt"
)
func DEMO_4_4_2main(){
	a := func(x, y int)(int){
		return x+y
	}
	result  := a(1,2)
	fmt.Println("result is" , result)	
}
