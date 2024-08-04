	
package main
import(
	"log"
)
type N int

func (n *N) pointer(){
	(*n)++
}
func test(){
	var a N = 1
	a.pointer()
	(&a).pointer()
	log.Println(a)
}

func main(){
	test()
}
