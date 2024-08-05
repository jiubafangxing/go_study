	
package main
import(
	"fmt"
)
type FuncSetting func() string

func (f FuncSetting) String()string{
	return f()
}

func test(){
	var t fmt.Stringer = FuncSetting(func ()string{
		return "hello world"
	})
	fmt.Println(t)
}

func main(){
	test()
}
