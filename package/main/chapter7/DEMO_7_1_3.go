	
package chapter7
import(
	"log"
)
type stringer interface{
	tostring()(string)
}
type DEMO_7_1_3tester interface{
	stringer
	DEMO_7_1_3test()
}

type N int
func (*N) DEMO_7_1_3test(){
	log.Println("DEMO_7_1_3test")
}

func (*N)tostring()(string){
	return "a"
}

func DEMO_7_1_3main(){
	var n N = 1
	n.DEMO_7_1_3test()
	log.Println(n.tostring())
}
