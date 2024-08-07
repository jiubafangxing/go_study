	
package chapter5
import(
	"log"
)
func DEMO_5_2_8test(a *[3]int){
	a[1] = 1
}

func DEMO_5_2_8main(){
	a := [...]int{2,2,2}
	DEMO_5_2_8test(&a)
	log.Println(a[1])
}
