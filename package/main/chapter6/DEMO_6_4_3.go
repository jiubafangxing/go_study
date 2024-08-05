	
package main
import(
	"log"
)
type N int

func (n N) changeAge(){
	log.Println(n)
	n++
}
func call(b func()){
	b()
}
func test(){
 	var a N = 1
	n := &a
	(*n)++
	call(n.changeAge)
	log.Println(*n)


}

func main(){
	test()
}
