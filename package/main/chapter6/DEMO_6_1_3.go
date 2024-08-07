	
package chapter6
import(
	"log"
)
type N int

func (n *N) pointer(){
	(*n)++
}
func DEMO_6_1_3test(){
	var a N = 1
	a.pointer()
	(&a).pointer()
	log.Println(a)
}

func DEMO_6_1_3main(){
	DEMO_6_1_3test()
}
