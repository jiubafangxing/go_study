package chapter4
import(
	"errors"
	"fmt"
)
func DEMO_4_3_3test(a, b int)(result int,err error){
    if b == 0{
	err = errors.New("division by zero")
	return 
    }
    result = a/b
    return

}

func DEMO_4_3_3main(){
	result,err := DEMO_4_3_3test(1,0)
	fmt.Println(result, err)
	result,err = DEMO_4_3_3test(4,2)
	fmt.Println(result, err)
}
