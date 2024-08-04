	
package main
import(
	"log"
)
type T struct{}
func (t *T) test(){
	log.Println("hi", t)
}
func test(){
	var t *T
	t.test()
	T{}.test()
}

func main(){
	test()
}
