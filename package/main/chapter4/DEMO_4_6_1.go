	
package chapter4
import(
	"log"
	"errors"
)
var errDivByZero = errors.New("division by zero")
func DEMO_4_6_1test(a , b int)(int, error){
	if b == 0{
		return 0, errDivByZero 
	}
	return a/b, nil
}

func DEMO_4_6_1main(){
	//result , err := DEMO_4_6_1test(1,2)
	result , err := DEMO_4_6_1test(1,0)
	if nil != err{
		log.Fatalln(err)
	}
	log.Println("result is ", result)
}
