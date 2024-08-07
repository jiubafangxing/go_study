	
package chapter5
import(
	"log"
)
func DEMO_5_2_7test(){
	a := [3]int{1,2,3}
	b := a
	b[1] = 4
	log.Println(a[1])
}

func DEMO_5_2_7main(){
	DEMO_5_2_7test()
}
