	
package chapter6
import(
	"fmt"
)

type N int

func (n N) ToString()(string){
	return fmt.Sprintf("%#x",n)

}


func DEMO_6_1_1test(){
	var n N = 1
	println(n.ToString)

}

func DEMO_6_1_1main(){
	DEMO_6_1_1test()
}
