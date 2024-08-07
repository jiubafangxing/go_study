	
package chapter5
import(
	"log"
)
func DEMO_5_2_20test(){
	b := make([]byte,3,5)
	n := copy(b,"abcd")
	log.Println(n,b)
}

func DEMO_5_2_20main(){
	DEMO_5_2_20test()
}
