	
package chapter6
import(
	"log"
)
type N int

func (n N) print1(){
	log.Println(n)
}
func DEMO_6_1_4test(){
	var n N = 1
	p := &n
	p.print1()
	p1 := &p
	p1.print1()

}

func DEMO_6_1_4main(){
	DEMO_6_1_4test()
}
