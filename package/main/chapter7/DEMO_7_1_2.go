	
package chapter7
import(
	"log"
)
func DEMO_7_1_2test(){
	var t1, t2 interface{}
	t1,t2 = 1,2
	t1,t2 = []int{},[]int{} 
	log.Println(t1 == t2)
}

func DEMO_7_1_2main(){
	DEMO_7_1_2test()
}
