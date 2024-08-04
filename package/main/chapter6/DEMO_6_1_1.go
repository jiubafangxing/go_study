	
package main
import(
	"fmt"
)

type N int

func (n N) ToString()(string){
	return fmt.Sprintf("%#x",n)

}


func test(){
	var n N = 1
	println(n.ToString)

}

func main(){
	test()
}
