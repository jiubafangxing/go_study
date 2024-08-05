	
package main
import(
	"log"
)
type N int
func (n N) testN(){
	log.Println("test.n:%p,%v\n", &n, n)
}
func test(){
	var n N = 1
	log.Println("test.n:%p,%v\n", &n, n)
	n++;
	test1 :=n.testN
	test1()
	log.Println("test.n:%p,%v\n", &n, n)
	test2 :=(&n).testN
	test2()
	log.Println("test.n:%p,%v\n", &n, n)
}

func main(){
	test()
}
