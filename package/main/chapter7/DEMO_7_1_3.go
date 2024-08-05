	
package main
import(
	"log"
)
type stringer interface{
	tostring()(string)
}
type tester interface{
	stringer
	test()
}

type N int
func (*N) test(){
	log.Println("test")
}

func (*N)tostring()(string){
	return "a"
}

func main(){
	var n N = 1
	n.test()
	log.Println(n.tostring())
}
