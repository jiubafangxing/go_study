	
package chapter5
import(
	"log"
)
func DEMO_5_1_8test(){
	var a []byte
	a = append(a,"hello"...)
	a = append(a,'a','b')
	log.Println(a[0])
}

func DEMO_5_1_8main(){
	DEMO_5_1_8test()
}
