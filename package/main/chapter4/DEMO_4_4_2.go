package main
import(
	"fmt"
)
func main(){
	a := func(x, y int)(int){
		return x+y
	}
	result  := a(1,2)
	fmt.Println("result is" , result)	
}
