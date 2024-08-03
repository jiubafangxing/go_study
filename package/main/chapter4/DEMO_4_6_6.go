	
package main
import(
	"runtime/debug"
)
func test(){
	panic("i am dead")
}

func main(){
	defer func (){
		if nil != recover(){
			debug.PrintStack()
		}
	}()
	test()
}
