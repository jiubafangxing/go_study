	
package chapter5
import(
	"log"
)
func DEMO_5_5_8test(){
	var b [100]struct{}
	log.Println(cap(b),len(b))
}

func DEMO_5_5_8main(){
	DEMO_5_5_8test()
}
