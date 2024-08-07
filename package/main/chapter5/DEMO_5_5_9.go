	
package chapter5
import(
	"log"
)
func DEMO_5_5_9test(){
	var a = struct{}{}
	var b = [3]struct{}{}
	var c = [0]int{}
	log.Println(a,b,c)
	log.Printf("%p", &a)
	log.Printf("%p", &b)
	log.Printf("%p", &c)


}

func DEMO_5_5_9main(){
	DEMO_5_5_9test()
}
