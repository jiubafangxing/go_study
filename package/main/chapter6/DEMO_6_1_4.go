	
package main
import(
	"log"
)
type N int

func (n N) print1(){
	log.Println(n)
}
func test(){
	var n N = 1
	p := &n
	p.print1()
	p1 := &p
	p1.print1()

}

func main(){
	test()
}
