package main
import(
	"fmt"
)
func test()(int, s string, e error){

	//can not compile
	s = "a"
	return 1,  "a" nil
}

func main(){
	a,b,err := test()
	fmt.Println(a,b,err)	
}
