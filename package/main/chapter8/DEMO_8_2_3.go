	
package chapter8
import(
	"log"
)
func DEMO_8_2_3test(){
	var a,b chan int = make (chan int, 3),make (chan int)
	log.Println(a == b)
	var c chan int
	log.Println(c == b)
	
}

func DEMO_8_2_3main(){
	DEMO_8_2_3test()
}
