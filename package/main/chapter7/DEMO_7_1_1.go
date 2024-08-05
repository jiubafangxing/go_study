	
package main
import(
	"log"
)
type tester interface{
	test()
	genStr()(string)
}
type N int
func (n N)test(){
	log.Println(n)
}
func (n *N)genStr()(string){
	return "adf"
}

func main(){
	var n N = 1
	n.test()
	s := n.genStr()
	log.Println(s)
	var t tester = &n
	t.test()
}
