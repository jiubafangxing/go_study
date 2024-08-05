	
package main
import(
	"log"
	"fmt"
)
type Ner interface{
	a()
	b(int)
	c(string)string
}
type N int

func (n N) a(){
	log.Println(n)
}

func (n *N) b(a N){
	*n = (*n) +a
}

func (n *N) c(b string)string{
	return fmt.Sprintf("%s-%d",b, *n)
}

func test(){
	var n N = 1
	var a N = 1
	n.a()
	n.b(a)
	log.Println(n)
	log.Println(n.c("hello"))
}

func main(){
	test()
}
