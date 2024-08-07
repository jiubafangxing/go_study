	
package chapter5
import(
	"log"
)
func DEMO_5_2_13test(){
	a := []int{1,2,3}
	a[1] = 3
	log.Println(a[1])
}

func DEMO_5_2_13main(){
	DEMO_5_2_13test()
}
