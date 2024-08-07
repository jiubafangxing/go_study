	
package chapter6
import(
	"log"
)
type N int
func (n N) DEMO_6_4_2testN(){
	log.Println("DEMO_6_4_2test.n:%p,%v\n", &n, n)
}
func DEMO_6_4_2test(){
	var n N = 1
	log.Println("DEMO_6_4_2test.n:%p,%v\n", &n, n)
	n++;
	DEMO_6_4_2test1 :=n.DEMO_6_4_2testN
	DEMO_6_4_2test1()
	log.Println("DEMO_6_4_2test.n:%p,%v\n", &n, n)
	DEMO_6_4_2test2 :=(&n).DEMO_6_4_2testN
	DEMO_6_4_2test2()
	log.Println("DEMO_6_4_2test.n:%p,%v\n", &n, n)
}

func DEMO_6_4_2main(){
	DEMO_6_4_2test()
}
