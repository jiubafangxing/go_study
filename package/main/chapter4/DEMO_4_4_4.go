package chapter4
import(
	"fmt"
)
func t( a func()(int)){
	fmt.Println("call func result is ", a())
}

func DEMO_4_4_4main(){
	a := func()(int){
		return 21
	}
	t(a)
}
