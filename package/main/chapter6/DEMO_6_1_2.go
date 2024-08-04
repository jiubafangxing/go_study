
package main
import(
	"log"
)
type N int
func (n N) value(){
	n++
	log.Printf("v:%p,%v\n",&n, n)
}

func (n *N) pointer(){
	(*n)++
	log.Printf("v:%p,%v\n",n, *n)
}
func test(){
	var n N  = 3
	n.value()
	log.Printf("v:%p,%v\n",&n, n)
	n.pointer()
	log.Printf("v:%p,%v\n",&n, n)
}

func main(){
	test()
}
