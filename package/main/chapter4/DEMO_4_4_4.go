package main
import(
	"fmt"
)
func t( a func()(int)){
	fmt.Println("call func result is ", a())
}

func main(){
	a := func()(int){
		return 21
	}
	t(a)
}
