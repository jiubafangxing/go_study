	
package chapter6
import(
	"log"
)
type T struct{}
func (t *T) DEMO_6_1_5test(){
	log.Println("hi", t)
}
func DEMO_6_1_5test(){
	var t *T
	t.DEMO_6_1_5test()
	T{}.DEMO_6_1_5test()
}

func DEMO_6_1_5main(){
	DEMO_6_1_5test()
}
