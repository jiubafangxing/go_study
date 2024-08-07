package chapter4
import(
	"fmt"
)
func DEMO_4_4_5main(){
	c := make(chan func(int, int)(int),2)
	c <- func(a, b int)(int){
		return a+b
	}
	fmt.Println("result is ",(<-c)(1,2))


}
