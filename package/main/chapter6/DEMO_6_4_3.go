	
package chapter6
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
func DEMO_6_4_3test(){
 	var a N = 1
	n := &a
	(*n)++
	call(n.changeAge)
	log.Println(*n)


}

func DEMO_6_4_3main(){
	DEMO_6_4_3test()
}
