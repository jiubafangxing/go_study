package main
import(
	"errors"
	"fmt"
)
func test(a, b int)(result int,err error){
    if b == 0{
	err = errors.New("division by zero")
	return 
    }
    result = a/b
    return

}

func main(){
	result,err := test(1,0)
	fmt.Println(result, err)
	result,err = test(4,2)
	fmt.Println(result, err)
}
