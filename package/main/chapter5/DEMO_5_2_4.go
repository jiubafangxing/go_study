	
package chapter5
import(
	"log"
)
func DEMO_5_2_4test(){
	a := [...][4]int{
		{1,2,3,4},
		{5,56,6,7},
	}
	log.Printf("#%v",a)
}

func DEMO_5_2_4main(){
	DEMO_5_2_4test()
}
