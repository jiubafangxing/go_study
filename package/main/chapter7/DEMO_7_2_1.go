	
package chapter7
import(
	"log"
)
func DEMO_7_2_1test(){
	var a interface{} = nil
	var b interface{}= (*int)(nil)
	log.Println(a == nil)
	log.Println(b == nil)
}

func DEMO_7_2_1main(){
	DEMO_7_2_1test()
}
