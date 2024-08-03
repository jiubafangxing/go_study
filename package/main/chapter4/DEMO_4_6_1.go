	
package main
import(
	"log"
	"errors"
)
var errDivByZero = errors.New("division by zero")
func test(a , b int)(int, error){
	if b == 0{
		return 0, errDivByZero 
	}
	return a/b, nil
}

func main(){
	//result , err := test(1,2)
	result , err := test(1,0)
	if nil != err{
		log.Fatalln(err)
	}
	log.Println("result is ", result)
}
