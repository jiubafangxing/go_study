	
package chapter4
import(
	"runtime/debug"
)
func DEMO_4_6_6test(){
	panic("i am dead")
}

func DEMO_4_6_6main(){
	defer func (){
		if nil != recover(){
			debug.PrintStack()
		}
	}()
	DEMO_4_6_6test()
}
