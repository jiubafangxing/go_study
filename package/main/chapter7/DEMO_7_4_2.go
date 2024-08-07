	
package chapter7
import(
	"fmt"
)
type FuncSetting func() string

func (f FuncSetting) String()string{
	return f()
}

func DEMO_7_4_2test(){
	var t fmt.Stringer = FuncSetting(func ()string{
		return "hello world"
	})
	fmt.Println(t)
}

func DEMO_7_4_2main(){
	DEMO_7_4_2test()
}
