	
package chapter5
import(
	"log"
)
func DEMO_5_1_7test(){
	s := "abcd"	
	a := []byte(s)
	a[0] = 'c'
	news :=string(a)
	log.Printf(news)
}

func DEMO_5_1_7main(){
	DEMO_5_1_7test()
}
