package chapter4
import(
	"fmt"
)
func DEMO_4_3_5test()(int, s string, e error){

	//can not compile
	s = "a"
	return 1,  "a" nil
}

func DEMO_4_3_5main(){
	a,b,err := DEMO_4_3_5test()
	fmt.Println(a,b,err)	
}
